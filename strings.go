package gostrings

import "regexp"

// Squishing the text
func Squish(text string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(text, " ")
}
