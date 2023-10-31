package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/acheong08/endless"
	gin "github.com/gin-gonic/gin"
	"github.com/xqdoo00o/funcaptcha"
)

var apiBreaker *funcaptcha.ApiBreaker

func main() {
	r := gin.Default()
	r.GET("/captcha/start", captchaStart)
	r.POST("/captcha/verify", captchaVerify)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	endless.ListenAndServe("127.0.0.1:"+port, r)
}

func captchaStart(c *gin.Context) {
	token, err := funcaptcha.GetOpenAITokenWithBx(4, `[{"key":"enhanced_fp","value":[{"key":"navigator_battery_charging","value":true}]},{"key":"fe","value":["DNT:1","L:zh-CN","D:24","PR:1","S:1920,1080","AS:1920,1080","TO:-480","SS:true","LS:true","IDB:true","B:false","ODB:true","CPUC:unknown","PK:Linux x86_64","CFP:11866 se","H:16","SWF:false"]}]`, "", "")
	if err == nil {
		c.JSON(200, gin.H{"token": token, "status": "success"})
		return
	}
	if err.Error() != "captcha required" {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	hex := "cd12da708fe6cbe6e068918c38de2ad9"
	session, err := funcaptcha.StartChallenge(token, hex)
	if err != nil {
		c.JSON(500, gin.H{"error": "unable to log requests"})
		return
	}
	apiBreaker, err = session.RequestChallenge(false)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to request challenge"})
		return
	}
	// Get form data (check if download_images is true)
	download_images := c.Query("download_images")
	var images []string
	if download_images == "true" {
		// Get Base64 encoded image
		images, err = funcaptcha.DownloadChallenge(session.ConciseChallenge.URLs, true)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to download images"})
			return
		}
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
	err := request.Session.SubmitAnswer([]int{request.Index}, false, apiBreaker)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Success
	c.JSON(200, gin.H{"status": "success"})
}
