package validate

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func validateUsername(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	// username can only contain letters, nums, underscores, and start only with lettters
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]+$`)
	return re.MatchString(value)
}

func validatePassword(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	var hasNum, hasUpper, hasSpecial bool

	if len(value) < 8 {
		return false
	}

	for _, c := range value {
		switch {
		case unicode.IsNumber(c):
			hasNum = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	return hasNum && hasUpper && hasSpecial
}
