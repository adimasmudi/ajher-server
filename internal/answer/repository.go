package answer

import "gorm.io/gorm"

type Repository interface {
	Save(answer Answer) (Answer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(answer Answer) (Answer, error) {
	err := r.db.Create(&answer).Error

	if err != nil {
		return answer, err
	}

	return answer, nil
}
