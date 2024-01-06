package participation

import (
	"ajher-server/internal/user"
	"time"
)

type Participation struct {
	ID        string    `json:"id"`
	UserId    int       `json:"user_id"`
	QuizId    string    `json:"quiz_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      user.User
}
