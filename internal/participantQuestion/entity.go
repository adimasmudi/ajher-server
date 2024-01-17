package participantQuestion

import (
	"ajher-server/internal/question"
	"time"
)

type ParticipantQuestion struct {
	ID              string            `json:"id" firestore:"-"`
	ParticipationId string            `json:"participation_id" firestore:"participantId"`
	QuestionId      string            `json:"question_id" firestore:"questionId"`
	Number          int64             `json:"number" firestore:"number"`
	CreatedAt       time.Time         `json:"created_at" firestore:"createdAt"`
	UpdatedAt       time.Time         `json:"updated_at" firestore:"updatedAt"`
	Question        question.Question `firestore:"question"`
}
