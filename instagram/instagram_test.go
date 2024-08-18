package instagram

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetFullBio(t *testing.T) {
	tests := []struct {
		name     string
		username string
		expected InstagramProfile
	}{
		{
			name:     "Valid User",
			username: "ariescarlas",
			expected: InstagramProfile{
				Title:     "Carla (@ariescarlas)",
				Bio:       "Halloo 🤔",
				Followers: 5,
				Following: 268,
				Post:      1,
			},
		},
		{
			name:     "User with no bio",
			username: "no_bio_user",
			expected: InstagramProfile{
				Title:     "",
				Bio:       "",
				Followers: 0,
				Following: 0,
				Post:      0,
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize colly collector
			c := Init()

			// Create a new Instagram module
			insta := NewInstagram(c)

			// Call GetFullBio with the test username
			profile := insta.GetFullBio(tt.username)

			// Check the results
			assert.Equal(t, tt.expected.Title, profile.Title, "Title should match")
			assert.Equal(t, tt.expected.Bio, profile.Bio, "Bio should match")
			assert.Equal(t, tt.expected.Followers, profile.Followers, "Followers count should match")
			assert.Equal(t, tt.expected.Following, profile.Following, "Following count should match")
			assert.Equal(t, tt.expected.Post, profile.Post, "Post count should match")
		})
	}
}
