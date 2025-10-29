package user

import (
	"context"
	"fmt"
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
	if _, err := s.repo.GetByUsername(ctx, payload.Username); err == nil {
		return nil, ErrUserAlreadyExists
	}

	// converting incoming `UserRegister` to `User` model
	newUser := payload.ToUser()

	// TODO: hash the password before creating

	createdUser, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	// converting `User` to `UserPublic`
	userPublic := createdUser.ToPublic()
	return userPublic, nil
}
