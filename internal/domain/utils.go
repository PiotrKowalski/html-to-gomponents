package domain

import (
	"regexp"
	"strings"
)

func removeBrackets(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "[", ""), "]", "")
}

func removeSpecialChars(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := re.ReplaceAllString(input, "")
	return result
}
