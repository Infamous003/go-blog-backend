package user

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID        int // json tags & validation isnt needed at the DB level
	Username  string
	Email     string
	Password  string
	Fname     string
	Lname     string
	CreatedAt time.Time
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

/*
GetAll returns a list of users ([]*UserPublic) or an error.
The argument required is a request context.
*/
func (r *Repository) GetAll(ctx context.Context) ([]*User, error) {
	query := `
		SELECT id, username, fname, lname, email, password, created_at
		FROM users
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Always close the rows

	var users []*User

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Fname,
			&u.Lname,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
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
		INSERT INTO users (username, fname, lname, email, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, fname, lname, email, password, created_at
	`

	var u User

	err := r.db.QueryRow(
		ctx,
		query,
		user.Username,
		user.Fname,
		user.Lname,
		user.Email,
		user.Password).
		Scan(
			&u.ID,
			&u.Username,
			&u.Fname,
			&u.Lname,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
		)

	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] CREATED: %v", u)
	log.Printf("[INFO] CREATED: %v", u.Username)
	return &u, nil
}
