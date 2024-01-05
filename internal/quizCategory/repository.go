package quizCategory

import "gorm.io/gorm"

type Repository interface {
	Save(quizCategory QuizCategory) (QuizCategory, error)
	Update(quizCategory QuizCategory) (QuizCategory, error)
	Delete(ID int) (QuizCategory, error)
	GetById(ID int) (QuizCategory, error)
	GetAll() ([]QuizCategory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(quizCategory QuizCategory) (QuizCategory, error) {
	err := r.db.Create(&quizCategory).Error

	if err != nil {
		return quizCategory, err
	}

	return quizCategory, nil
}

func (r *repository) Update(quizCategory QuizCategory) (QuizCategory, error) {
	err := r.db.Save(&quizCategory).Error

	if err != nil {
		return quizCategory, err
	}

	return quizCategory, nil
}

func (r *repository) Delete(ID int) (QuizCategory, error) {
	var quizCategory QuizCategory
	err := r.db.Where("id=?", ID).Delete(&quizCategory).Error

	if err != nil {
		return quizCategory, err
	}

	return quizCategory, nil
}

func (r *repository) GetById(ID int) (QuizCategory, error) {
	var quizCategory QuizCategory
	err := r.db.Where("id=?", ID).Find(&quizCategory).Error

	if err != nil {
		return quizCategory, err
	}

	return quizCategory, nil
}

func (r *repository) GetAll() ([]QuizCategory, error) {
	var quizCategories []QuizCategory

	error := r.db.Find(&quizCategories).Error

	if error != nil {
		return quizCategories, error
	}

	return quizCategories, nil
}
