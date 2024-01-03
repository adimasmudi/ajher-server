package ranking

import (
	"time"

	"gorm.io/gorm"
)

type Ranking struct {
	gorm.Model
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
