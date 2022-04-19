package user

import "time"

type Users struct {
	Id       int
	Username string
	Email    string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}