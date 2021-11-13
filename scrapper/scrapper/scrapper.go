package scrapper

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// URL for getting the most recent posts.
var baseURL = "https://kr.indeed.com/jobs?q=ruby&sort=date"

type jobPost struct {
	title   string
	name    string
	address string
	link    string
}

func Scrape(keyword string) {
	changeBaseURL(keyword)
	channel := make(chan []jobPost)
	extractedPosts := []jobPost{}
	pages := getPageLength(baseURL)

	// Navigate the HTML.
	for i := 0; i < pages; i++ {
		go getPageContent(i, channel)
	}

	// Show all the extracted information.
	for i := 0; i < pages; i++ {
		extractedPosts = append(extractedPosts, <-channel...)
	}

	// Save the extracted information in a 'csv' file.
	saveExtractedJobs(keyword, extractedPosts)
}

func changeBaseURL(keyword string) {
	baseURL = "https://kr.indeed.com/jobs?q=" + keyword + "&sort=date"
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

func getPageContent(page int, channel chan<- []jobPost) {
	extractChannel := make(chan jobPost)

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
	extractedPosts := []jobPost{}
	cards := doc.Find(".tapItem")
	cards.Each(func(i int, s *goquery.Selection) {
		go extractJobPost(s, viewURL, extractChannel)
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

func extractJobPost(card *goquery.Selection, viewURL string, extractChannel chan<- jobPost) {
	id, _ := card.Attr("data-jk")
	post := jobPost{
		title:   card.Find(".jobTitle > span").Text(),
		name:    card.Find(".companyName").Text(),
		address: card.Find(".companyLocation").Text(),
		link:    viewURL + id,
	}
	extractChannel <- post
}

func saveExtractedJobs(keyword string, extractedJobs []jobPost) {
	writeJobChannel := make(chan []string)
	file, createErr := os.Create(keyword + "_jobs.csv")
	checkError(createErr)

	writer := csv.NewWriter(file)
	defer writer.Flush()
	defer file.Close() // For making os.Remove() work in Win10 + Chrome environment.

	headers := []string{
		"Title",
		"Name",
		"Address",
		"Link",
	}

	writeErr := writer.Write(headers)
	checkError(writeErr)

	for _, job := range extractedJobs {
		go writeJob(job, writeJobChannel)
	}

	for i := 0; i < len(extractedJobs); i++ {
		jobWriteErr := writer.Write(<-writeJobChannel)
		checkError(jobWriteErr)
	}
}

func writeJob(job jobPost, writeJobChannel chan<- []string) {
	postedJob := []string{job.title, job.name, job.address, job.link}
	writeJobChannel <- postedJob
}
