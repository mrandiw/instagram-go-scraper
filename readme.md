[![Build Status](https://github.com/MrAndiw/instagram-go-scraper/actions/workflows/go.yml/badge.svg)](https://github.com/MrAndiw/instagram-go-scraper/actions)
<img src="https://camo.githubusercontent.com/33cfae3047a121e7811c7a54e7b6ef4029c9db941f3d180a176069220c878954/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f676f2d2532333030414444382e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d676f266c6f676f436f6c6f723d7768697465" alt="Go" data-canonical-src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&amp;logo=go&amp;logoColor=white" style="max-width: 100%;">
<img src="https://camo.githubusercontent.com/7f5701ed50f919cf2352cd028b5b2dc974b5e643fe4d78ad826eb9e74551157f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f496e7374616772616d2d2532334534343035462e7376673f7374796c653d666f722d7468652d6261646765266c6f676f3d496e7374616772616d266c6f676f436f6c6f723d7768697465" alt="Instagram" data-canonical-src="https://img.shields.io/badge/Instagram-%23E4405F.svg?style=for-the-badge&amp;logo=Instagram&amp;logoColor=white" style="max-width: 100%;">

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
