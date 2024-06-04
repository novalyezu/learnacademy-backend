package user

import "time"

type User struct {
	ID        string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
