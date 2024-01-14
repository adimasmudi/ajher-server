package participantQuestion

import (
	"ajher-server/internal/question"
	"time"
)

type ParticipantQuestion struct {
	ID              string    `json:"id"`
	ParticipationId string    `json:"participation_id"`
	QuestionId      string    `json:"question_id"`
	Number          int       `json:"number"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Question        question.Question
}
