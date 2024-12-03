package actio

import (
	"fmt"
	"regexp"
)

// RegexpTest checks if the given text matches the specified regular expression pattern.
//
// Example Usage:
//
//	println(RegexpTest("hello123", "^[a-z]+[0-9]+$")) // true
func RegexpTest(text, pattern string) bool {
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
