package quizCategory

import (
	"time"
)

type QuizCategory struct {
	ID           string    `json:"id" firestore:"-"`
	CategoryName string    `json:"category_name" firestore:"categoryName"`
	Icon         string    `json:"icon" firestore:"icon"`
	Description  string    `json:"description" firestore:"description"`
	CreatedAt    time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt    time.Time `json:"updated_at" firestore:"updatedAt"`
}
