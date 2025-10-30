package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

// A sepcial function that Go runs automatically when the package is loaded
func init() {
	v = validator.New()

	v.RegisterValidation("username", validateUsername)
	v.RegisterValidation("password", validatePassword)
}

// Struct validates a struct using Go's validator, kinda like a wrapper
func Struct(s any) error {
	return v.Struct(s)
}

// returns a slice of errors occured during validation
func ErrorMessages(err error) []string {
	var errs []string
	if err == nil {
		return errs
	}

	for _, e := range err.(validator.ValidationErrors) {
		msg := fmt.Sprintf("Field '%s' failed validation rule '%s'", e.Field(), e.Tag())
		errs = append(errs, msg)
	}
	return errs
}
