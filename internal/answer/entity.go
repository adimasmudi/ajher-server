package answer

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	ID                  string    `json:"id"`
	UserId              int       `json:"user_id"`
	QuestionId          string    `json:"question_id"`
	Grade               float64   `json:"grade"`
	Label               string    `json:"label"`
	Answer              string    `json:"answer"`
	AnswerDuration      int       `json:"answer_duration"`
	Status              string    `json:"status"`
	GeneratedSuggestion string    `json:"generated_suggestion"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
