package otp

import (
	"time"
)

type Otp struct {
	ID         int       `json:"id"`
	UserId     int       `json:"userId"`
	Otpcode    string    `json:"otpCode"`
	Status     string    `json:"status"`
	ValidUntil time.Time `json:"validUntil"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
