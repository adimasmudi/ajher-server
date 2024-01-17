package participantRanking

import (
	"time"
)

type ParticipantRanking struct {
	ID              string    `json:"id" firestore:"-"`
	ParticipationId string    `json:"participation_id" firestore:"participantId"`
	RankingId       int64     `json:"ranking_id" firestore:"rankingId"`
	Grade           float64   `json:"grade" firestore:"grade"`
	Point           int64     `json:"point" firestore:"point"`
	Note            string    `json:"note" firestore:"note"`
	Position        int64     `json:"position" firestore:"position"`
	CreatedAt       time.Time `json:"created_at" firestore:"createdAt"`
	UpdatedAt       time.Time `json:"updated_at" firestore:"updatedAt"`
}
