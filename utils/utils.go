package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Infamous003/go-blog-backend/types"
)

func FromJSON(r io.Reader, payload any) error {
	if payload == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r).Decode(payload)
}

func ToJSON(w io.Writer, data any) error {
	if data == nil {
		return fmt.Errorf("missing data to encode")
	}
	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	e := types.ErrorResponse{
		Status:  status,
		Message: message,
	}
	return ToJSON(w, e)
}
