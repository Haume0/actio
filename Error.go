package actio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ErrorCodes map[string]ErrorCode
var Languages = []string{"tr", "en"}

type ErrorCode struct {
	Message map[string]string `json:"message"`
	Status  int               `json:"status"`
}
type errorResponse struct {
	Error string `json:"error"`
}

// FindCode finds the error code in the error codes map
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

// Error : Makes default error response with message and status code
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

// Example JSON Object
/*
"AUTH.BAD_REQUEST": {
		    "Message": {
		      "tr": "Geçersiz istek",
		      "en": "Bad request",
		      "ru": "Плохой запрос",
		      "es": "Solicitud incorrecta"
		    },
		    "Status": 400
		  }
*/
func ErrorInit(errorCodesPath string) {
	file, err := os.Open(errorCodesPath)
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
