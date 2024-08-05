package instagram

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse the string and return the struct
func parseFollowers(input string) (InstagramProfile, error) {
	parts := strings.Split(input, ", ")
	if len(parts) != 3 {
		return InstagramProfile{}, fmt.Errorf("input string does not have exactly 3 parts")
	}

	followers, err := parseNumber(parts[0])
	if err != nil {
		return InstagramProfile{}, err
	}

	following, err := parseNumber(parts[1])
	if err != nil {
		return InstagramProfile{}, err
	}

	posts, err := parseNumber(parts[2])
	if err != nil {
		return InstagramProfile{}, err
	}

	return InstagramProfile{Followers: followers, Following: following, Post: posts}, nil
}

// Helper function to parse the numbers and handle 'K' suffix
func parseNumber(part string) (int, error) {
	numberStr := strings.Fields(part)[0]
	numberStr = strings.ReplaceAll(numberStr, ",", "")
	if strings.HasSuffix(numberStr, "K") {
		numberStr = strings.TrimSuffix(numberStr, "K")
		value, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return 0, err
		}
		return int(value * 1000), nil
	}
	return strconv.Atoi(numberStr)
}
