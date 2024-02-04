package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCollyScraper_NotAllowedDomain(t *testing.T) {
	//Capture printed messages
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	// Run the scraper with a domain not allowed (Wikipedia de)
	url := "https://de.wikipedia.org/wiki/Data_Science"
	runCollyScraper(url)

	// Close the write end of the pipe and read the output
	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Check the output
	expectedOutput := "Error: Forbidden domain"
	assert.Contains(t, buf.String(), expectedOutput)
}

func TestMainFunction(t *testing.T) {
	//Run program
	main()
}
