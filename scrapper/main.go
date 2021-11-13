package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/tyomhk2015/go_basics/scrapper/scrapper"
)

func main() {
	// Create a server /w Echo
	e := echo.New()
	// Set routes for the server
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
}

// When users access to the root, show them a HTML template.
func handleHome(c echo.Context) error {
	return c.File("home.html")
}

// When users clicks 'Scrape' button, the application gets keyword(s) from the form.
const fileName string = "_jobs.csv"

func handleScrape(c echo.Context) error {
	keyword := strings.ToLower(strings.TrimSpace(c.FormValue("keyword")))

	// Erase "jobs.csv" created by scrapper.go
	defer os.Remove(keyword + fileName)

	scrapper.Scrape(keyword)

	fmt.Println(keyword + fileName)
	return c.Attachment(keyword+fileName, keyword+fileName)
}
