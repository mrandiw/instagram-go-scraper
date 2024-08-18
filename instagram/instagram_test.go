package instagram

import (
	"fmt"
	"testing"
)

func TestGetFullBio(t *testing.T) {

	// Initialize colly collector
	c := Init()

	// Create a new Instagram module
	insta := NewInstagram(c)

	// Call GetFullBio with the test username
	profile := insta.GetFullBio("ariescarlas")
	fmt.Println(profile)

	// Validate the scraped data
	expectedProfile := InstagramProfile{
		Title:        "Carla (@ariescarlas)",
		Bio:          "Halloo 🤔",
		Followers:    5,
		Following:    268,
		Post:         1,
		PhotoProfile: "https://scontent.cdninstagram.com/v/t51.2885-19/162875510_451501566051191_8601471211802227301_n.jpg?stp=dst-jpg_p100x100&_nc_cat=110&ccb=1-7&_nc_sid=fcb8ef&_nc_ohc=eHGO8MFoAGoQ7kNvgHAdjfi&_nc_ht=scontent.cdninstagram.com&oh=00_AYB8qNnkq1oZAtD5NHkKDDR6VGW67WM4VBB6eG8bA0XwnA&oe=66C728C1",
	}

	if profile != expectedProfile {
		t.Errorf("expected %+v, got %+v", expectedProfile, profile)
	}
}
