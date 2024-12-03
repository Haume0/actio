package actio

import (
	"fmt"
	"log"
)

// Unwrap logs error and returns value
// If err is not nil, logs the error and exits. Otherwise, returns value.
func Unwrap[T any](v T, err error) T {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return v
}

// Ok prints the error if it is not nil.
func Ok(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
