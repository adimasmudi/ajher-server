package questionImage

import (
	"time"

	"gorm.io/gorm"
)

type QuestionImage struct {
	gorm.Model
	ID         string    `json:"id"`
	QuestionId string    `json:"question_id"`
	ImagePath  string    `json:"image_path"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
