package participantQuestion

import (
	"time"

	"gorm.io/gorm"
)

type ParticipantQuestion struct {
	gorm.Model
	ID              string    `json:"id"`
	ParticipationId string    `json:"participation_id"`
	QuestionId      string    `json:"question_id"`
	Number          int       `json:"number"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
