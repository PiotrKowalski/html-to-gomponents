package domain

import (
	"regexp"
)

func removeSpecialChars(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := re.ReplaceAllString(input, "")
	return result
}
