package profile

import (
	"fmt"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

func GetProfileExample() {
	// Instantiate default collector
	ig := instagram.Init()

	Instagram := instagram.NewInstagram(ig)

	User := Instagram.GetFullBio("ariescarlas")
	fmt.Println("Name : ", User.Title)
	fmt.Println("Bio :", User.Bio)
	fmt.Println("Followers :", User.Followers)
	fmt.Println("Following :", User.Following)
	fmt.Println("Post :", User.Post)
	fmt.Println("Photo Profile :", User.PhotoProfile)
}
