package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	channel := make(chan bool)
	extractedPosts := []jobPost{}
	pages := getPageLength(baseURL)

	// Navigate the HTML.
	for i := 0; i < pages; i++ {
		go getPageContent(i, &extractedPosts, channel)
	}

	// Show all the extracted information.
	for i := 0; i < pages; i++ {
		<-channel
	}

	// Save the extracted information in a 'csv' file.
	saveExtractedJobs(extractedPosts)
	fmt.Println("Scrapping is done!")
}

var (
	// URL for getting the most recent posts.
	baseURL = "https://kr.indeed.com/jobs?q=ruby&sort=date"
)

type jobPost struct {
	title   string
	name    string
	address string
	link    string
}

func getPageLength(url string) int {
	// Request the HTML page.
	res, err := http.Get(url)
	checkError(err)
	checkResponse(res)

	// Prevent memory leak.
	defer preventMemoryLeak(res)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	// Find the pagination element.
	totalPages := 0
	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		// Get the total number of pages in one HTML page.
		totalPages = s.Find("a").Length()
	})
	return totalPages
}

func getPageContent(page int, extractedPosts *[]jobPost, channel chan<- bool) {
	pageURL := baseURL + "&start=" + strconv.Itoa(10*page)
	// Load the HTML document
	res, err := http.Get(pageURL)
	checkError(err)
	checkResponse(res)

	// Prevent memory leak.
	defer preventMemoryLeak(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	// Return all the job posts from current page.
	viewURL := "https://kr.indeed.com/viewjob?jk="
	cards := doc.Find(".tapItem")
	cards.Each(func(i int, s *goquery.Selection) {
		post := extractJobPost(s, viewURL)
		postSlice := []jobPost{*post}
		*extractedPosts = append(*extractedPosts, postSlice...)
	})

	channel <- true
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

func extractJobPost(card *goquery.Selection, viewURL string) *jobPost {
	id, _ := card.Attr("data-jk")
	post := jobPost{
		title:   card.Find(".jobTitle > span").Text(),
		name:    card.Find(".companyName").Text(),
		address: card.Find(".companyLocation").Text(),
		link:    viewURL + id,
	}
	return &post
}

func saveExtractedJobs(extractedJobs []jobPost) {
	file, createErr := os.Create("jobs.csv")
	checkError(createErr)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"Title",
		"Name",
		"Address",
		"Link",
	}

	writeErr := writer.Write(headers)
	checkError(writeErr)

	for _, job := range extractedJobs {
		jobSlice := []string{job.title, job.name, job.address, job.link}
		jobWriteErr := writer.Write(jobSlice)
		checkError(jobWriteErr)
	}
}
