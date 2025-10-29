package user

import (
	"context"
	"fmt"
	"log"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// RegisterUser creates a new user if the username isn't already taken.
// ErrUserAlreadyExists is one of the expected errors.
// If there are no errors, it returns the created user (*UserPublic)
func (s *Service) RegisterUser(ctx context.Context, payload *UserRegister) (*UserPublic, error) {
	// CHecking if the user already exists or not
	_, err := s.repo.GetByUsername(ctx, payload.Username)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	log.Printf("[INFO] User exists: %v", err)
	log.Printf("[INFO] Username: %v", payload.Username)
	// converting incoming `UserRegister` to `User` model
	newUser := &User{
		Username: payload.Username,
		Fname:    payload.Fname,
		Lname:    payload.Lname,
		Email:    payload.Email,
		Password: payload.Password,
	}

	// TODO: hash the password before creating

	createdUser, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	// converting `User` to `UserPublic`
	userPublic := &UserPublic{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Fname:    createdUser.Fname,
		Lname:    createdUser.Lname,
		Email:    createdUser.Email,
	}

	return userPublic, nil
}
