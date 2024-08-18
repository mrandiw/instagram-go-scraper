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
	expectedBio := "Tech Whiz 🚀 • Let's Innovate Together • #SoftwareEngineer"

	// Validate results
	if User.Title != expectedTitle {
		t.Errorf("expected Title %v, got %v", expectedTitle, User.Title)
	}
	if User.Bio != expectedBio {
		t.Errorf("expected Followers %v, got %v", expectedBio, User.Bio)
	}
}
