package instagram

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

type (
	ModuleInstagram interface {
		visit(username string)
		getTitle()
		getBio()
		GetFullBio(username string) InstagramProfile
	}

	moduleInstagram struct {
		c                *colly.Collector
		InstagramProfile InstagramProfile
	}

	// InstagramProfile holds the scraped data
	InstagramProfile struct {
		Title     string
		Bio       string
		Followers int
		Following int
		Post      int
	}
)

var (
	Title = ""
)

func NewInstagram(c *colly.Collector) ModuleInstagram {
	return &moduleInstagram{
		c: c,
	}
}

func Init() *colly.Collector {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("instagram.com", "www.instagram.com"),
	)

	return c
}

func (module *moduleInstagram) visit(username string) {
	// Replace 'username' with the actual Instagram username
	profileURL := fmt.Sprintf("https://www.instagram.com/%s/", username)
	err := module.c.Visit(profileURL)
	if err != nil {
		log.Fatal(err)
	}
}

func (module *moduleInstagram) getTitle() {

	// Find and print title bio
	module.c.OnHTML("meta[property=\"og:title\"]", func(e *colly.HTMLElement) {

		title := e.Attr("content")

		end := strings.Index(title, "•")
		if end != -1 {
			result := title[:end]

			module.InstagramProfile.Title = strings.TrimSpace(result)
		} else {
			module.InstagramProfile.Title = ""
		}
	})
}

func (module *moduleInstagram) getBio() {

	// Find and print bio and other details
	module.c.OnHTML("meta[name=\"description\"]", func(e *colly.HTMLElement) {

		bio := e.Attr("content")

		// parse bio desc
		re := regexp.MustCompile(`"([^"]+)"`)
		match := re.FindStringSubmatch(bio)

		bioDesc := ""
		if len(match) > 1 {
			bioDesc = match[1]
		}

		end := strings.Index(bio, "-")
		if end != -1 {
			result := bio[:end]

			// parse follower
			bioFollower, err := parseFollowers(strings.TrimSpace(result))
			if err != nil {
				log.Fatal(err)
			}

			// assign value
			module.InstagramProfile.Bio = bioDesc
			module.InstagramProfile.Followers = bioFollower.Followers
			module.InstagramProfile.Following = bioFollower.Following
			module.InstagramProfile.Post = bioFollower.Post
		} else {
			module.InstagramProfile.Followers = 0
			module.InstagramProfile.Following = 0
			module.InstagramProfile.Post = 0
		}
	})

}

func (module *moduleInstagram) GetFullBio(username string) InstagramProfile {

	// get title
	module.getTitle()

	// get bio
	module.getBio()

	// Error handling
	module.c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Replace 'username' with the actual Instagram username
	module.visit(username)

	profile := InstagramProfile{
		Title:     module.InstagramProfile.Title,
		Bio:       module.InstagramProfile.Bio,
		Followers: module.InstagramProfile.Followers,
		Following: module.InstagramProfile.Following,
		Post:      module.InstagramProfile.Post,
	}

	return profile
}
