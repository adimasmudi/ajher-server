package user

import (
	"time"
)

type User struct {
	ID        int       `json:"ID"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	Picture   string    `json:"picture"`
	Gender    string    `json:"gender"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
