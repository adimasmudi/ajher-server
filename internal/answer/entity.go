package answer

import (
	"ajher-server/internal/question"
	"time"
)

type Answer struct {
	ID                  string    `json:"id" firestore:"-"`
	UserID              string    `json:"user_id" firestore:"userId"`
	QuestionID          string    `json:"question_id" firestore:"questionId"`
	Grade               float64   `json:"grade" firestore:"grade"`
	Label               string    `json:"label" firestore:"label"`
	AnswerText          string    `json:"answer_text" firestore:"answerText"`
	AnswerDuration      int64     `json:"answer_duration" firestore:"answerDuration"`
	Status              string    `json:"status" firestore:"status"`
	GeneratedSuggestion string    `json:"generated_suggestion" firestore:"generatedSuggestion"`
	CreatedAt           time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt           time.Time `json:"updated_at" firestore:"updatedAt"`
}

type AnswerWithQuestion struct {
	Answer
	Question question.Question `json:"question"`
}

func (a *Answer) BeforeCreate() {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Answer) BeforeUpdate() {
	a.UpdatedAt = time.Now()
}
