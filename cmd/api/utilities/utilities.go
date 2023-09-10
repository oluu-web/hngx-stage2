package utilities

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes a JSON response to the provided http.ResponseWriter.
func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	// Create a map to wrap the data with the specified key.
	wrapper := make(map[string]interface{})
	wrapper[wrap] = data

	// Marshal the wrapper map to JSON.
	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	// Set the Content-Type header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP status code and write the JSON response.
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// ErrorJSON writes an error JSON response to the provided http.ResponseWriter.
func ErrorJSON(w http.ResponseWriter, err error) {
	// Define a JSONError struct to encapsulate the error message.
	type JSONError struct {
		Message string `json:"message"`
	}

	// Create a JSONError instance with the error message.
	theError := JSONError{
		Message: err.Error(),
	}

	// Write the error JSON response with a HTTP status code of Bad Request.
	WriteJSON(w, http.StatusBadRequest, theError, "error")
}
