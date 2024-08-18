package instagram

import (
	"testing"
)

func TestParseFollowers(t *testing.T) {
	tests := []struct {
		input       string
		expected    InstagramProfile
		expectError bool
	}{
		{
			input: "1,200 followers, 300 following, 150 posts",
			expected: InstagramProfile{
				Followers: 1200,
				Following: 300,
				Post:      150,
			},
			expectError: false,
		},
		{
			input: "2K followers, 1.5K following, 200 posts",
			expected: InstagramProfile{
				Followers: 2000,
				Following: 1500,
				Post:      200,
			},
			expectError: false,
		},
		{
			input: "1.2M followers, 500 following, 20 posts",
			expected: InstagramProfile{
				Followers: 1200000,
				Following: 500,
				Post:      20,
			},
			expectError: false,
		},
		{
			input:       "500 followers, 300 following",
			expected:    InstagramProfile{},
			expectError: true,
		},
		{
			input:       "invalid followers, 300 following, 150 posts",
			expected:    InstagramProfile{},
			expectError: true,
		},
	}

	for _, tt := range tests {
		result, err := parseFollowers(tt.input)
		if tt.expectError {
			if err == nil {
				t.Errorf("expected error but got none for input: %s", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %+v, got %+v", tt.expected, result)
			}
		}
	}
}
