# Instagram Go Scraper

`instagram-go-scraper` is a Go library for scraping Instagram user profiles. It allows you to fetch details such as bio, followers count, following count, and the number of posts from Instagram profiles.

## Installation

To use this library, you need to have Go installed on your system. You can install the library using the following command:

```bash
go get github.com/MrAndiw/instagram-go-scraper
```

## Example Code

This code bellow for initialize the library

```bash
// Initialize the Instagram client
Initialize the Instagram client
ig := instagram.Init()
```

Create a new Instagram scraper instance

```bash
// Create a new Instagram scraper instance
Instagram := instagram.NewInstagram(ig)
```

// Fetch the full bio information for a specific Instagram user

```bash
// Fetch the full bio information for a specific Instagram user
User := Instagram.GetFullBio("mrandiw")
```

## Run Test

```bash
go test -v ./...
```
