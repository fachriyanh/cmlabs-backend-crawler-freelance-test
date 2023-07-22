package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	websiteList := map[string]string{
		"cmlabs":   "https://cmlabs.co",
		"sequence": "https://sequence.day",
		"kompas":   "https://tekno.kompas.com/apps-os",
	}

	// Create a channel to synchronize Goroutines
	done := make(chan struct{})

	// Iterate through the websiteList and start a Goroutine for each entry
	for title, url := range websiteList {
		go craw(title, url, done)
	}

	// Wait for all Goroutines to finish
	for range websiteList {
		<-done
	}
}

func craw(title string, url string, done chan<- struct{}) {
	// Create a new headless browser context with chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set up a timeout for the headless browser context
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Navigate to the SPA, SSR, or PWA website using chromedp
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Fatalf("Error navigating to the page: %v", err)
	}

	// Wait for the page to load (adjust the sleep duration based on the website)
	time.Sleep(2 * time.Second)

	// Get the fully rendered HTML content using chromedp
	var htmlContent string
	err = chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery))
	if err != nil {
		log.Fatalf("Error extracting fully rendered HTML: %v", err)
	}

	// Specify the folder path where the result should be saved
	outputFolder := "output"
	outputName := title + "_output.html"

	// Save the fully rendered HTML to a new file
	outputFile := filepath.Join(outputFolder, outputName)
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(htmlContent)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	fmt.Printf("Result saved to %s\n", outputFile)

	// Signal the completion of the Goroutine
	done <- struct{}{}
}
