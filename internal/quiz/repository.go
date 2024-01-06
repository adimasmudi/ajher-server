package quiz

import "gorm.io/gorm"

type Repository interface {
	Save(quiz Quiz) (Quiz, error)
	GetById(id string) (Quiz, error)
	GetByCode(code string) (Quiz, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(quiz Quiz) (Quiz, error) {
	err := r.db.Create(&quiz).Error

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}

func (r *repository) GetById(id string) (Quiz, error) {
	var quiz Quiz
	err := r.db.Preload("Question").Preload("QuizCategory").Where("id=?", id).Find(&quiz).Error

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}

func (r *repository) GetByCode(code string) (Quiz, error) {
	var quiz Quiz
	err := r.db.Preload("Question").Preload("QuizCategory").Where("code=?", code).Find(&quiz).Error

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}
