package participation

import "gorm.io/gorm"

type Repository interface {
	Save(partcipation Participation) (Participation, error)
	GetByQuizId(quizId string) (Participation, error)
	GetByUserId(userId int) (Participation, error)
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

func (r *repository) GetByQuizId(quizId string) (Participation, error) {
	var participation Participation
	err := r.db.Preload("User").Where("quiz_id=?", quizId).Find(&participation).Error

	if err != nil {
		return participation, err
	}

	return participation, nil
}

func (r *repository) GetByUserId(userId int) (Participation, error) {
	var participation Participation
	err := r.db.Preload("User").Where("user_id=?", userId).Find(&participation).Error

	if err != nil {
		return participation, err
	}

	return participation, nil
}
