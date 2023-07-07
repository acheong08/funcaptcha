package main

import (
	"log"

	"github.com/acheong08/funcaptcha"
)

func main() {
	for {
		token, hex, err := funcaptcha.GetOpenAIToken()
		log.Println(token)

		if err == nil {
			return
		}
		log.Printf("error getting token: %v\n", err)
		// Start a challenge
		session, err := funcaptcha.StartChallenge(token, hex)
		if err != nil {
			log.Fatalf("error starting challenge: %v\n", err)
		}
		log.Println("Challenge started!")

		err = session.RequestChallenge(false)
		if err != nil {
			log.Fatalf("error requesting challenge: %v\n", err)
		}
		log.Println(session.ConciseChallenge)

		err = session.SubmitAnswer(2)
		if err != nil {
			log.Printf("error submitting answer: %v\n", err)
			continue
		} else {
			log.Println("Answer submitted!")
			break
		}
	}
}
