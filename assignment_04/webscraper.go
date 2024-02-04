package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

// WikiPage creates a struct representing details of the Wikipedia pages
type WikiPage struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {

	// Visit specified URLs and scrape using colly
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	var pageData []WikiPage

	//Loop through Wikipedia pages and append page details to pageData
	for _, url := range urls {
		data := runCollyScraper(url)
		pageData = append(pageData, data)
	}

	// Save the extracted data to a JSON file
	jsonData, err := json.MarshalIndent(pageData, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	filePath := "items.json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Printf("Data saved to %s\n", filePath)
	fmt.Println("\nScraping completed\n")
}

func runCollyScraper(url string) WikiPage {
	c := colly.NewCollector(
		//Visit only domains: en.wikipedia.org
		colly.AllowedDomains("en.wikipedia.org"),
	)

	var WikiPage WikiPage

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		WikiPage.Title = e.Text
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Extract only <p> text content without HTML markup
		WikiPage.Text = e.DOM.Find("p").Text()
	})

	// Start the scraping process
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	WikiPage.URL = url
	return WikiPage
}
