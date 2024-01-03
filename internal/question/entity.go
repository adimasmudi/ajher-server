package question

import (
	"ajher-server/internal/answer"
	"ajher-server/internal/questionImage"
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID              string    `json:"id"`
	QuizId          string    `json:"quiz_id"`
	Question        string    `json:"question"`
	ReferenceAnswer string    `json:"reference_answer"`
	GradePercentage float64   `json:"grade_percentage"`
	Status          string    `json:"status"`
	Duration        int       `json:"duration"`
	Point           float64   `json:"point"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	QuestionImage   []questionImage.QuestionImage
	Answer          []answer.Answer
}
