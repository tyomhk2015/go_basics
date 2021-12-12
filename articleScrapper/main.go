package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	start := time.Now()
	channel := make(chan []articlePost)
	extractedPosts := []articlePost{}

	// Navigate the HTML.
	for i := 0; i < len(category); i++ {
		go getPageContent(category[i], channel)
	}

	// // Show all the extracted information.
	for i := 0; i < len(category); i++ {
		extractedPosts = append(extractedPosts, <-channel...)
	}

	// // Save the extracted information in a 'csv' file.
	saveExtractedArticles(extractedPosts)
	t := time.Now()
	fmt.Println("The scrapping is done!", "\nOperation Time:", t.Sub(start))
}

const (
	domain string = "https://b.hatena.ne.jp"
)

var (
	// URL for getting the most recent posts.
	baseURL  = "https://b.hatena.ne.jp/hotentry/"
	category = []string{"all", "general", "social", "economics", "life", "knowledge", "it", "fun", "entertainment", "game"}
)

// If the fields' initial letter don't start with uppercase,
// then the json parser won't able to get any information about them.
type articlePost struct {
	Category    string   `json:"category"`
	Title       string   `json:"title"`
	Users       string   `json:"users"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	PostedDate  string   `json:"postedDate"`
	Tags        []string `json:"tags"`
}

type tag struct {
	Tag string `json:"tag"`
}

func getPageContent(category string, channel chan<- []articlePost) {
	extractChannel := make(chan articlePost)
	pageURL := baseURL + category
	// Load the HTML document
	res, err := http.Get(pageURL)
	checkError(err)
	checkResponse(res)

	// Prevent memory leak.
	defer preventMemoryLeak(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	// Return all the article posts from current page.
	extractedPosts := []articlePost{}
	cards := doc.Find(".entrylist-contents-main")
	cards.Each(func(i int, s *goquery.Selection) {
		go extractArticlePost(s, category, extractChannel)
		extractedPosts = append(extractedPosts, <-extractChannel)
	})
	channel <- extractedPosts
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Request Failed.", err)
	}
}

func checkResponse(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatal("Request Failed.", res.StatusCode)
	}
}

func preventMemoryLeak(res *http.Response) {
	res.Body.Close()
}

func extractArticlePost(card *goquery.Selection, category string, extractChannel chan<- articlePost) {
	var tags = []string{}

	retrievedTags := card.Find(".entrylist-contents-tags").Children()
	for i := 0; i < retrievedTags.Length(); i++ {
		tags = append(tags, retrievedTags.Eq(i).Text())
	}

	link, _ := card.Find(".entrylist-contents-users").Children().Attr("href")

	post := articlePost{
		Category:    category,
		Title:       card.Find(".entrylist-contents-title").Children().Text(),
		Users:       card.Find(".entrylist-contents-users").Find("span").Text(),
		Link:        domain + link,
		Description: card.Find(".entrylist-contents-body").Find("p").Text(),
		PostedDate:  card.Find(".entrylist-contents-date").Text(),
		Tags:        tags,
	}
	extractChannel <- post
}

func saveExtractedArticles(extractedArticles []articlePost) {
	jsonData, _ := json.MarshalIndent(extractedArticles, "", " ")
	ioutil.WriteFile("articles.json", jsonData, 0666)
}
