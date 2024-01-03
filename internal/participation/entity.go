package participation

import (
	"time"

	"gorm.io/gorm"
)

type Participation struct {
	gorm.Model
	ID        string    `json:"id"`
	UserId    int       `json:"user_id"`
	QuizId    string    `json:"quiz_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
