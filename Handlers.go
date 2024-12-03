package actio

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Unwrap : handles error and returns value
func Unwrap[T any](v T, err error) T {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return v
}

// Ok : handles error and prints it
func Ok(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Jsonify : returns json from struct
func Jsonify[T any](data T) string {
	res, err := json.Marshal(data)
	if err != nil {
		return `{
			"error": "error marshalling data"
			}`
	}
	return string(res)
}

// PlainText : returns plain text from struct
func PlainText(data map[string]string) string {
	//loop through struct key-value pairs
	var res string = ""
	for _, value := range data {
		// value|value|value
		if res == "" {
			res = value
			continue
		}
		res = strings.Join([]string{res, value}, "|")
	}
	return res
}
