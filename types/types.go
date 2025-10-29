package types

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"` // a short description of the error
}
