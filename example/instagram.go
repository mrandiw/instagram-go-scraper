package main

import (
	"fmt"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

func main() {
	// Instantiate default collector
	ig := instagram.Init()

	Instagram := instagram.NewInstagram(ig)

	User := Instagram.GetFullBio("mrandiw")
	fmt.Println(User.Title)
	fmt.Println(User.Bio)
	fmt.Println(User.Followers)
	fmt.Println(User.Following)
	fmt.Println(User.Post)
}
