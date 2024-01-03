package questionImage

import (
	"time"
)

type QuestionImage struct {
	ID         string    `json:"id"`
	QuestionId string    `json:"question_id"`
	ImagePath  string    `json:"image_path"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
