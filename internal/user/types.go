package user

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserRegister struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Fname    string `json:"fname" validate:"required,max=32"`
	Lname    string `json:"lname" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,email,max=64"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserPublic struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
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
	return validate.Struct(u)
}

// Converts User model to UserPublic. Excluds fields like password
func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
		ID:       u.ID,
		Username: u.Username,
		Fname:    u.Fname,
		Lname:    u.Lname,
		Email:    u.Email,
	}
}

// Converts UserRegister to User model, which represents user in DB
func (u *UserRegister) ToUser() *User {
	return &User{
		Username: u.Username,
		Fname:    u.Fname,
		Lname:    u.Lname,
		Email:    u.Email,
		Password: u.Password,
	}
}
