package instagram

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetFullBio(t *testing.T) {

	// Initialize colly collector
	c := Init()

	// Create a new Instagram module
	insta := NewInstagram(c)

	// Call GetFullBio with the test username
	profile := insta.GetFullBio("ariescarlas")

	// Now we can check the results
	assert.Equal(t, "Carla (@ariescarlas)", profile.Title, "Title should match")
	assert.Equal(t, "Halloo 🤔", profile.Bio, "Bio should match")
	assert.Equal(t, 5, profile.Followers, "Followers count should match")
	assert.Equal(t, 268, profile.Following, "Following count should match")
	assert.Equal(t, 1, profile.Post, "Post count should match")
}
