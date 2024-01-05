package quizCategory

import (
	"time"
)

type QuizCategory struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	Icon         string    `json:"icon"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
