package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// pages := getPageLength(baseURL)
	first_page := getPageContent(1)
	fmt.Println("Current Page", first_page)
}

// ..&start=10
var baseURL = "https://kr.indeed.com/jobs?q=ruby&sort=date"

func getPageLength(url string) int {
	// Request the HTML page.
	res, err := http.Get(url)
	checkError(err)
	checkResponse(res)

	// Prevent memory leak.
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	checkResponse(res)

	// Find the pagination element.
	totalPages := 0
	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		// Get the total number of pages in one HTML page.
		intValue, _ := strconv.Atoi(s.Find("b").Text())
		totalPages = intValue
	})
	return totalPages
}

func getPageContent(page int) int {
	pageURL := baseURL + "&start=" + strconv.Itoa(10*page)
	first_page := getPageLength(pageURL)
	return first_page
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
