package utils

import (
	"fmt"
	"log"
	"net/http"
)

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
