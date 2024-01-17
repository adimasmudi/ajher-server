package quiz

import (
	"ajher-server/internal/question"
	"ajher-server/internal/quizCategory"
	"time"
)

type Quiz struct {
	ID             string                    `json:"id" firestore:"-"`
	QuizCategoryId string                    `json:"quiz_category_id" firestore:"quizCategoryId"`
	Title          string                    `json:"title" firestore:"title"`
	Duration       int                       `json:"duration" firestore:"duration"`
	EndAt          time.Time                 `json:"end_at" firestore:"endAt"`
	Code           string                    `json:"code" firestore:"code"`
	Status         string                    `json:"status" firestore:"status"`
	Description    string                    `json:"description" firestore:"description"`
	CreatedAt      time.Time                 `json:"created_at" firestore:"createdAt"`
	UpdatedAt      time.Time                 `json:"updated_at" firestore:"updatedAt"`
	QuizCategory   quizCategory.QuizCategory `firestore:"quizCategory"`
	Question       []question.Question       `firestore:"question"`
}
