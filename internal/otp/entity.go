package otp

import (
	"time"
)

type Otp struct {
	ID         int       `json:"id"`
	UserId     string    `json:"userId"`
	Otpcode    string    `json:"otpCode"`
	Status     string    `json:"status"`
	ValidUntil string    `json:"validUntil"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
