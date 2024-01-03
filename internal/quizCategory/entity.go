package quizCategory

import (
	"time"

	"gorm.io/gorm"
)

type QuizCategory struct {
	gorm.Model
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	Icon         string    `json:"icon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
