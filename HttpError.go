package actio

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorCode struct {
	Message map[string]string `json:"message"`
	Status  int               `json:"status"`
}
type errorResponse struct {
	Error string `json:"error"`
}

// ErrorCodes: Error Code List
var ErrorCodes map[string]ErrorCode

// Languages: Available Languages
var Languages = []string{"tr", "en"}

func findCode(language, code string) (ErrorCode, bool) {
	if errorCode, ok := ErrorCodes[code]; ok {
		if message, ok := errorCode.Message[language]; ok {
			return ErrorCode{
				Message: map[string]string{language: message},
				Status:  errorCode.Status,
			}, true
		}
	}
	return ErrorCode{}, false
}

// Error makes the default HTTP error response prettier and localized.
//
// Example Usage:
//
//	actio.Error(w, "User not found", http.StatusNotFound)
//	actio.Error(w, "User not found", http.StatusNotFound, "tr")
//
// Migration Notes:
//   - Old: http.Error(w, "User not found", http.StatusNotFound)
//   - New:	actio.Error(w, "User not found", http.StatusNotFound, "tr")
func Error(w http.ResponseWriter, message string, status int, lang ...string) {
	var language string = "en"

	if len(lang) > 0 {
		language = lang[0]
		errorCode, ok := findCode(language, message)

		if ok {
			message = errorCode.Message[language]
			status = errorCode.Status
		}
	}

	// Creating JSON response
	res, err := json.Marshal(errorResponse{Error: message})
	if err != nil {
		http.Error(w, "error marshalling data", http.StatusInternalServerError)
		return
	}

	// Setting header
	w.Header().Set("Content-Type", "application/json")

	// Writing response
	w.WriteHeader(status)
	w.Write(res)
}

// ErrorInit initializes the error codes and available languages.
// This is an optional operation if localized error messages are to be used.
//
// Example Usage:
//
//	actio.ErrorInit(errorCodes, []string{"tr", "en"})
func ErrorInit(errorCodes map[string]ErrorCode, availableLanguages []string) {
	ErrorCodes = errorCodes
	Languages = availableLanguages
	fmt.Println("✅ Error codes loaded!")
}
