package participation

import "gorm.io/gorm"

type Repository interface {
	Save(partcipation Participation) (Participation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(participation Participation) (Participation, error) {
	err := r.db.Create(&participation).Error

	if err != nil {
		return participation, err
	}

	return participation, nil
}
