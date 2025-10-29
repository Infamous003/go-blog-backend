package user

import "time"

type User struct {
	ID        int // json tags & validation isnt needed at the DB level
	Username  string
	Email     string
	Password  string
	Fname     string
	Lname     string
	CreatedAt time.Time
}
