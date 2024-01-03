package leaderboard

import (
	"ajher-server/internal/ranking"
	"ajher-server/internal/user"
	"time"

	"gorm.io/gorm"
)

type Leaderboard struct {
	gorm.Model
	ID         int       `json:"id"`
	UserId     int       `json:"user_id"`
	RankingId  int       `json:"ranking_id"`
	TotalPoint int       `json:"total_point"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Ranking    ranking.Ranking
	User       []user.User
}
