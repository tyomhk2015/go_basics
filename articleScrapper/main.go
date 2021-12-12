package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

type articlePost struct {
	title       string
	users       string
	link        string
	description string
	postedDate  string
	tags        []string
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
		go extractArticlePost(s, extractChannel)
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

func extractArticlePost(card *goquery.Selection, extractChannel chan<- articlePost) {
	var tags = []string{}

	retrievedTags := card.Find(".entrylist-contents-tags").Children()
	for i := 0; i < retrievedTags.Length(); i++ {
		if i == retrievedTags.Length()-1 {
			tags = append(tags, retrievedTags.Eq(i).Text())
		} else {
			tags = append(tags, retrievedTags.Eq(i).Text()+"/")
		}
	}

	link, _ := card.Find(".entrylist-contents-users").Children().Attr("href")

	post := articlePost{
		title:       card.Find(".entrylist-contents-title").Children().Text(),
		users:       card.Find(".entrylist-contents-users").Find("span").Text(),
		link:        domain + link,
		description: card.Find(".entrylist-contents-body").Find("p").Text(),
		postedDate:  card.Find(".entrylist-contents-date").Text(),
		tags:        tags,
	}
	extractChannel <- post
}

func saveExtractedArticles(extractedArticles []articlePost) {
	writeArticleChannel := make(chan []string)
	file, createErr := os.Create("articles.csv")
	checkError(createErr)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"title",
		"users",
		"link",
		"description",
		"postedDate",
		"tags",
	}

	writeErr := writer.Write(headers)
	checkError(writeErr)

	for _, article := range extractedArticles {
		go writeArticle(article, writeArticleChannel)
	}

	for i := 0; i < len(extractedArticles); i++ {
		articleWriteErr := writer.Write(<-writeArticleChannel)
		checkError(articleWriteErr)
	}
}

func writeArticle(article articlePost, writeArticleChannel chan<- []string) {
	postedArticle := []string{article.title, article.users, article.link, article.description, article.postedDate, fmt.Sprint(article.tags)}
	writeArticleChannel <- postedArticle
}
