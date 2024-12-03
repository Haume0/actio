package actio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ErrorCode struct {
	Message map[string]string `json:"message"`
	Status  int               `json:"status"`
}
type errorResponse struct {
	Error string `json:"error"`
}

// FindErrorCode belirli bir hata kodunu arar ve döner
func FindErrorCode(language, code string) (ErrorCode, bool) {
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

// Error : makes default http error response prettier
func Error(w http.ResponseWriter, message string, status int, lang ...string) {
	var language string = "en"
	if len(lang) > 0 {
		language = lang[0]
	}
	errorCode, ok := FindErrorCode(language, message)
	if ok {
		message = errorCode.Message[language]
		status = errorCode.Status
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

var ErrorCodes map[string]ErrorCode

func ErrorInit() {
	file, err := os.Open("error_codes.json")
	if err != nil {
		panic("Error opening error codes file: " + err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ErrorCodes)
	if err != nil {
		panic("Error decoding error codes: " + err.Error())
	}
	fmt.Println("✅ Error codes loaded!")
}
