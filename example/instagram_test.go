package main

import (
	"testing"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

func TestGetFullBio(t *testing.T) {
	// Mock or initialize the Instagram client
	ig := instagram.Init()
	Instagram := instagram.NewInstagram(ig)

	// Call the method to test
	User := Instagram.GetFullBio("mrandiw")

	// Define expected values
	expectedTitle := "𝑨𝒏𝒅𝒊 𝑾𝒊𝒃𝒐𝒘𝒐 (@mrandiw)"
	expectedFollowers := 12000
	expectedFollowing := 1053
	expectedPosts := 624

	// Validate results
	if User.Title != expectedTitle {
		t.Errorf("expected Title %v, got %v", expectedTitle, User.Title)
	}
	if User.Followers != expectedFollowers {
		t.Errorf("expected Followers %v, got %v", expectedFollowers, User.Followers)
	}
	if User.Following != expectedFollowing {
		t.Errorf("expected Following %v, got %v", expectedFollowing, User.Following)
	}
	if User.Post != expectedPosts {
		t.Errorf("expected Post %v, got %v", expectedPosts, User.Post)
	}
}
