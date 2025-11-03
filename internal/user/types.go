package user

import (
	"regexp"
	"strings"

	"github.com/Infamous003/go-blog-backend/internal/validator"
)

type UserRegister struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Fname    string `json:"fname" validate:"required,max=32"`
	Lname    string `json:"lname" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,email,max=64"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type UserPublic struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
}

type UserUpdate struct {
	Username string `json:"username" validate:"required,min=8,max=32"`
	Fname    string `json:"fname" validate:"required,max=32"`
	Lname    string `json:"lname" validate:"required,max=32"`
	Email    string `json:"email" validate:"required,email,max=64"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

var (
	UsernameRX = regexp.MustCompile(`^[A-Za-z\d_.@]+$`)
	PasswordRX = regexp.MustCompile(`^[A-Za-z\d_.@]+$`)
	EmailRX    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func ValidateUser(v *validator.Validator, u *UserRegister) {
	// validating username
	v.Check(u.Username != "", "username", "must be provided")
	v.Check(len(u.Username) >= 8 && len(u.Username) <= 32, "username", "must be between 8 and 32 characters long")
	v.Check(validator.Matches(u.Username, UsernameRX), "username", "must contain at least 1 uppercase letter, 1 number, and a special character (_, ., @)")

	// validating firstname
	u.Fname = strings.TrimSpace(u.Fname)
	v.Check(u.Fname != "", "fname", "must be provided")
	v.Check(len(u.Fname) <= 32, "fname", "must not be more than 32 characters long")

	// validating lastname
	u.Lname = strings.TrimSpace(u.Lname)
	v.Check(u.Lname != "", "lname", "must be provided")
	v.Check(len(u.Lname) <= 32, "lname", "must not be more than 32 characters long")

	// validating email
	u.Email = strings.TrimSpace(u.Email)
	v.Check(u.Email != "", "email", "must be provided")
	v.Check(len(u.Email) <= 64, "email", "must not be more than 64 characters long")
	v.Check(validator.Matches(u.Email, EmailRX), "email", "must be a valid email address")

	// validating password
	v.Check(u.Password != "", "password", "must be provided")
	v.Check(len(u.Password) >= 8 && len(u.Password) <= 64, "password", "must be between 8 and 64 characters long")
	v.Check(validator.Matches(u.Password, PasswordRX), "password", "must contain at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character (_, ., @)")

	v.Check(strings.ContainsAny(u.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		"password", "must contain at least 1 uppercase letter")
	v.Check(strings.ContainsAny(u.Password, "abcdefghijklmnopqrstuvwxyz"),
		"password", "must contain at least 1 lowercase letter")
	v.Check(strings.ContainsAny(u.Password, "0123456789"),
		"password", "must contain at least 1 number")
	v.Check(strings.ContainsAny(u.Password, "_.@"),
		"password", "must contain at least 1 special character (_, ., @)")

}
