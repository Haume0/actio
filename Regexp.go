package actio

import (
	"fmt"
	"regexp"
)

// TestRegex checks if the given text matches the provided regex pattern.
// It returns true if the text matches the pattern, otherwise false.
func RegexTest(text, pattern string) bool {
	// Compile the regular expression pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		// Print error if regex compilation fails
		fmt.Println("Error compiling regex:", err)
		return false
	}

	// Check if the text matches the compiled regex pattern
	return re.MatchString(text)
}
