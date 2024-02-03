package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

// PageInfo represents information extracted from a webpage
type PageInfo struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Function to list all files and directories in a given directory
func listAll(currentDirectory string) {
	filepath.Walk(currentDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(path)
		return nil
	})
}

func main() {
	// Make directory for storing complete HTML code for web page
	pageDirname := "wikipages"
	if _, err := os.Stat(pageDirname); os.IsNotExist(err) {
		os.Mkdir(pageDirname, os.ModePerm)
	}

	// Examine the directory structure
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	listAll(currentDirectory)

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

	var pageData []PageInfo

	for _, url := range urls {
		data := runCollyScraper(url)
		pageData = append(pageData, data)
	}

	// Save the extracted data to a JSON file
	saveToJSON(pageData)

	fmt.Println("\nScraping completed\n")
}

func runCollyScraper(url string) PageInfo {
	c := colly.NewCollector()

	// Set up callbacks for scraping logic
	var pageInfo PageInfo

	c.OnHTML("title", func(e *colly.HTMLElement) {
		pageInfo.Title = e.Text
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Extract text content without HTML markup
		pageInfo.Text = e.DOM.Find("p").Text() // Example: Extract text only from <p> tags
	})

	// Set up callback to save HTML to a file
	c.OnHTML("html", func(e *colly.HTMLElement) {
		filePath := fmt.Sprintf("wikipages/%s.html", e.Request.URL.Hostname())
		err := e.Response.Save(filePath)
		if err != nil {
			fmt.Println("Error saving HTML:", err)
		}
		fmt.Printf("HTML saved to %s\n", filePath)
	})

	// Start the scraping process
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}

	pageInfo.URL = url
	return pageInfo
}

func saveToJSON(data []PageInfo) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
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
}
