package quiz

import (
	"ajher-server/internal/question"
	"ajher-server/internal/quizCategory"
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID             string    `json:"id"`
	QuizCategoryId int       `json:"quiz_category_id"`
	Title          string    `json:"title"`
	Duration       int       `json:"duration"`
	EndAt          time.Time `json:"end_at"`
	Code           string    `json:"code"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	QuizCategory   quizCategory.QuizCategory
	Question       question.Question
}
