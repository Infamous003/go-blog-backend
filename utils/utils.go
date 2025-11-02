package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func FromJSON(r io.Reader, payload any) error {
	if payload == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r).Decode(payload)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}
	return nil
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
