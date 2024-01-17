package otp

import (
	"time"
)

type Otp struct {
	ID         string    `json:"id" firestore:"-"`
	UserId     string    `json:"userId" firestore:"userId"`
	Otpcode    string    `json:"otpCode" firestore:"otpCode"`
	Status     string    `json:"status" firestore:"status"`
	ValidUntil time.Time `json:"validUntil" firestore:"validUntil"`
	CreatedAt  time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt  time.Time `json:"updated_at" firestore:"updatedAt"`
}
