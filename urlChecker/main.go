package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://cover-corp.com/",
		"https://hrmos.co/pages/cover-corp/jobs/711-1",
		"https://www.youtube.com/user/tanigox",
		"https://www.sandbox.game/en/",
		"https://www.epicgames.com/fortnite/en-US/home",
		"https://studio.zepeto.me/",
		"https://www.binance.com/en",
	}

	channel := make(chan requestResult)

	for _, url := range urls {
		go checkUrls(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		requestResult := <-channel
		fmt.Println(requestResult.url, requestResult.result)
	}
}

var (
	errRequestFailed = errors.New("❌")
)

type requestResult struct {
	url    string
	result string
}

func checkUrls(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	result := "✔️"
	if err != nil || resp.StatusCode >= 400 {
		result = "❌"
	}
	c <- requestResult{url: url, result: result}
}

func channel_main() {
	// variable := make(channel data_type_for_communication)
	channel := make(chan bool)
	foods := [3]string{"Sushi", "Ekiben", "Natto"}
	for _, food := range foods {
		// go routine for concurrency
		go eatFood(food, channel)
	}
}

// The second argument: Takes a channel, and a bool for communicating with main func
func eatFood(food string, c chan bool) {
	fmt.Println(food)
	// Send bool to channel for communicating w/ the main func.
	c <- true
}
