package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func getIDFromURL(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return -1, err
	}
	return id, nil
}
