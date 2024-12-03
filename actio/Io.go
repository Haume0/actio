package actio

import (
	"encoding/json"
	"io"
	"os"
)

// GetFileAsStruct : reads file and returns it as struct
func GetFileAsStruct[T any](filename string) (T, error) {
	var result T
	//read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return result, err
	}
	//parse file
	var file T
	err = json.Unmarshal(data, &file)
	if err != nil {
		return result, err
	}
	result = file
	return result, nil
}

// WriteUpload : Saves formdata to file
func WriteUplaod(filename string, data io.Reader) (string, error) {
	fileBytes, err := io.ReadAll(data)
	if err != nil {
		return "", err
	}
	//write file
	err = os.WriteFile(filename, fileBytes, 0644)
	if err != nil {
		return "", err
	}
	return filename, nil
}
