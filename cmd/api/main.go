package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/acheong08/endless"
	"github.com/acheong08/funcaptcha"
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/captcha/start", captchaStart)
	r.POST("/captcha/verify", captchaVerify)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	endless.ListenAndServe(":"+port, r)
}

func captchaStart(c *gin.Context) {
	var images []string
	var token, hex string
	var err error
	var session *funcaptcha.Session
	for i := 0; i < 20; i++ {
		token, hex, err = funcaptcha.GetOpenAIToken()
		if err == nil {
			c.JSON(200, gin.H{"token": token, "status": "success"})
			return
		}
		if err.Error() != "captcha required" {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		session, err = funcaptcha.StartChallenge(token, hex)
		if err != nil {
			c.JSON(500, gin.H{"error": "unable to log requests"})
			return
		}
		err = session.RequestChallenge(false)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to request challenge"})
			return
		}
		err = session.SubmitAnswer(2)
		if err == nil {
			c.JSON(200, gin.H{"token": token, "status": "success"})
			return
		}
		log.Println("Retrying...")
	}

	// Download the images
	images, err = funcaptcha.DownloadChallenge(session.ConciseChallenge.URLs, true)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to download images"})
	}
	c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{"token": token, "session": session, "status": "captcha", "images": images})
}

func captchaVerify(c *gin.Context) {
	type submissionRequest struct {
		Index   int                `json:"index"`
		Session funcaptcha.Session `json:"session"`
	}
	var request submissionRequest
	// Map the request body to the submissionRequest struct
	if c.Request.Body != nil {
		err := json.NewDecoder(c.Request.Body).Decode(&request)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(400, gin.H{"error": "request body not provided"})
		return
	}
	// Verify the captcha
	err := request.Session.SubmitAnswer(request.Index)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Success
	c.JSON(200, gin.H{"status": "success"})
}
