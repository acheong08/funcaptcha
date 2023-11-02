package main

import (
	"fmt"
	"log"

	"github.com/acheong08/funcaptcha"
)

func main() {
	token := "155179050361eafc5.5354759702|r=us-west-2|meta=3|metabgclr=transparent|metaiconclr=%23757575|guitextcolor=%23000000|pk=0A1D34FC-659D-4E23-B17B-694DCFCF6A6C|at=40|ag=101|cdn_url=https%3A%2F%2Ftcr9i.chat.openai.com%2Fcdn%2Ffc|lurl=https%3A%2F%2Faudio-us-west-2.arkoselabs.com|surl=https%3A%2F%2Ftcr9i.chat.openai.com|smurl=https%3A%2F%2Ftcr9i.chat.openai.com%2Fcdn%2Ffc%2Fassets%2Fstyle-manager"
	hex := "fbfc14b0d793c6ef8359e0e4b4a91f67"
	// Start a challenge
	session, err := funcaptcha.StartChallenge(token, hex)
	if err != nil {
		log.Fatalf("error starting challenge: %v\n", err)
	}
	log.Println("Challenge started!")

	apiBreaker, err := session.RequestChallenge(false)
	if err != nil {
		log.Fatalf("error requesting challenge: %v\n", err)
	}
	log.Println(session.ConciseChallenge)
	log.Println("Downloading challenge")
	_, err = funcaptcha.DownloadChallenge(session.ConciseChallenge.URLs, false)
	if err != nil {
		log.Fatalf("error downloading challenge: %v\n", err)
	}
	log.Println("Challenge downloaded!")
	// User input here
	fmt.Println("Please enter the index of the image based on the following instructions:")
	fmt.Println(session.ConciseChallenge.Instructions)

	ints := make([]int, len(session.ConciseChallenge.URLs))
	fmt.Println("Enter the integers:")
	for i := 0; i < len(ints); i++ {
		fmt.Scan(&ints[i])
	}
	log.Println(ints)
	err = session.SubmitAnswer(ints, false, apiBreaker)
	if err != nil {
		log.Fatalf("error submitting answer: %v\n", err)
	}
}
