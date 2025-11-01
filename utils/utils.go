package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FromJSON(r io.Reader, payload any) error {
	if payload == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	dataBytes = append(dataBytes, '\n')

	for k, v := range headers {
		// can use .Add(k, v) as well, but we want to overwrite any headers
		// add, adds a new key-val pair, even if it exists
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)

	return nil
}

func WriteError(w http.ResponseWriter, status int, message string) {
	data := map[string]string{
		"error": message,
	}
	err := WriteJSON(w, status, data, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	WriteError(w, http.StatusNotFound, message)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	log.Printf("[ERROR] %v", err)
	WriteError(w, http.StatusInternalServerError, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	WriteError(w, http.StatusMethodNotAllowed, message)
}
