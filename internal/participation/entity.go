package participation

import (
	"ajher-server/internal/user"
	"time"
)

type Participation struct {
	ID        string    `json:"id" firestore:"-"`
	UserId    string    `json:"user_id" firestore:"userId"`
	QuizId    string    `json:"quiz_id" firestore:"quizId"`
	Status    string    `json:"status" firestore:"status"`
	CreatedAt time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updatedAt"`
	User      user.User `firestore:"user"`
}
