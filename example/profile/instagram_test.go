package main

import (
	"fmt"
	"testing"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

func TestMain(t *testing.T) {
	// Test case 1: Initialize the Instagram client
	ig := instagram.Init()
	if ig == nil {
		t.Errorf("Failed to initialize the Instagram client")
	}

	// Test case 2: Create a new Instagram scraper instance
	Instagram := instagram.NewInstagram(ig)
	if Instagram == nil {
		t.Errorf("Failed to create a new Instagram scraper instance")
	}

	// Test case 3: Fetch the full bio information for a specific Instagram user
	User := Instagram.GetFullBio("ariescarlas")
	if User.Title == "" {
		t.Errorf("Failed to fetch the title of the user profile")
	}
	if User.Bio == "" {
		t.Errorf("Failed to fetch the bio of the user profile")
	}
	if User.Followers == 0 {
		t.Errorf("Failed to fetch the followers count of the user profile")
	}
	if User.Following == 0 {
		t.Errorf("Failed to fetch the following count of the user profile")
	}
	if User.Post == 0 {
		t.Errorf("Failed to fetch the post count of the user profile")
	}
	if User.PhotoProfile == "" {
		t.Errorf("Failed to fetch the photo profile of the user profile")
	}

	// Test case 4: Print the correct information from the fetched user profile
	expectedOutput := "Name : Carla (@ariescarlas)\nBio : Halloo 🤔\nFollowers : 5\nFollowing : 268\nPost : 1\n"
	actualOutput := fmt.Sprintf("Name : %s\nBio : %s\nFollowers : %d\nFollowing : %d\nPost : %d\n", User.Title, User.Bio, User.Followers, User.Following, User.Post)
	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, actualOutput)
	}
}
