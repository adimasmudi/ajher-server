package user

import (
	"time"
)

type User struct {
	ID        string    `json:"ID" firestore:"-"`
	FullName  string    `json:"fullName" firestore:"fullname"`
	Email     string    `json:"email" firestore:"email"`
	Password  string    `json:"password" firestore:"password"`
	Username  string    `json:"username" firestore:"username"`
	Picture   string    `json:"picture" firestore:"picture"`
	Gender    string    `json:"gender" firestore:"gender"`
	LastLogin time.Time `json:"last_login" firestore:"lastLogin"`
	CreatedAt time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updatedAt"`
}
