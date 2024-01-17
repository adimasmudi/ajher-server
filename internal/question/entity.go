package question

import (
	"ajher-server/internal/questionImage"
	"time"
)

type Question struct {
	ID              string                        `json:"id" firestore:"-"`
	QuizId          string                        `json:"quiz_id" firestore:"quizId"`
	Question        string                        `json:"question" firestore:"question"`
	ReferenceAnswer string                        `json:"reference_answer" firestore:"referenceAnswer"`
	GradePercentage float64                       `json:"grade_percentage" firestore:"gradePercentage"`
	Status          string                        `json:"status" firestore:"status"`
	Duration        int64                         `json:"duration" firestore:"duration"`
	Point           float64                       `json:"point" firestore:"point"`
	CreatedAt       time.Time                     `json:"created_at" firestore:"createdAt"`
	UpdatedAt       time.Time                     `json:"updated_at" firestore:"updatedAt"`
	QuestionImage   []questionImage.QuestionImage `firestore:"questionImages"`
}
