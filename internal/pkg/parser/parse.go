package parser

import "strings"

// Parse parses string input
func Parse(input string) []string {
	return strings.Split(input, "")
}
