package user

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserRegister struct {
	Fname    string `json:"fname" validate:"required,max=32"`
	Lname    string `json:"lname" validate:"required,max=32"`
	Username string `json:"username" validate:"required,min=4,max=32"`
	Email    string `json:"email" validate:"required,email,max=64"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserPublic struct {
	ID       int    `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}

type ErrorResonse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (u *UserRegister) Validate() error {
	// validate.RegisterValidation("username", validateUsername)
	return validate.Struct(u)
}

// func validateUsername(fl validator.FieldLevel) bool {
// 	if fl.Field().String() == "admin" {
// 		return false
// 	}

// 	switch fl.Field().String() {
// 	case "admin":
// 		return false
// 	case "nigga":
// 		return false
// 	case "fuck":
// 		return false
// 	}

// 	return true
// }
