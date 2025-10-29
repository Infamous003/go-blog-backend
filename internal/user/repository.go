package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Our data access layer
// All the functions that let you access the User data lie here

type Repository struct {
	db *pgxpool.Pool
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

/*
GetAll returns a list of users ([]*UserPublic) or an error.
The argument required is a request context.
*/
func (r *Repository) GetAll(ctx context.Context) ([]*User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, fname, lname, username, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Always close the rows

	var users []*User

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Fname,
			&u.Lname,
			&u.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

/*
GetByUsername returns a user (*User) or an error.
It takes in a context and a username as args.
*/
func (r *Repository) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `
		SELECT id, username, fname, lname, email, password, created_at
		FROM users
		WHERE username=$1
	`

	// *UserPublic would be nil pointer. Access to it(&user.ID) would panic
	var user User
	err := r.db.QueryRow(
		ctx,
		query,
		username).
		Scan(
			&user.ID,
			&user.Username,
			&user.Fname,
			&user.Lname,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Create(ctx context.Context, user *User) (*User, error) {
	query := `
		INSERT INTO users (fname, lname, username, password, email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, fname, lname, email, password, created_at
	`

	var u User

	err := r.db.QueryRow(
		ctx,
		query,
		user.Fname,
		user.Lname,
		user.Username,
		user.Password,
		user.Email).
		Scan(
			&u.ID,
			&user.Username,
			&u.Fname,
			&u.Lname,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
