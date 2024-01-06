package participantQuestion

import "gorm.io/gorm"

type Repository interface {
	Save(question []ParticipantQuestion) ([]ParticipantQuestion, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(question []ParticipantQuestion) ([]ParticipantQuestion, error) {
	err := r.db.Create(&question).Error

	if err != nil {
		return question, err
	}

	return question, nil
}
