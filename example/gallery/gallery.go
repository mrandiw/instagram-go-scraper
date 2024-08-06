package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type InstagramPhoto struct {
	URL   string
	Alt   string
	Thumb string
}

func main() {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the Instagram profile page
	username := "mrandiw"
	profileURL := fmt.Sprintf("https://www.instagram.com/%s/", username)

	var photos []InstagramPhoto

	// Run the chromedp tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate(profileURL),
		chromedp.Sleep(5*time.Second), // Wait for the page to load
		chromedp.EvaluateAsDevTools(`Array.from(document.querySelectorAll('article img')).map(img => ({url: img.src, alt: img.alt, thumb: img.srcset.split(",")[0].split(" ")[0]}))`, &photos),
	)

	if err != nil {
		log.Fatal(err)
	}

	// Print the fetched photo data
	for _, photo := range photos {
		fmt.Printf("Photo URL: %s\nAlt Text: %s\nThumbnail: %s\n\n", photo.URL, photo.Alt, photo.Thumb)
	}
}
