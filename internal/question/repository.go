package question

import "gorm.io/gorm"

type Repository interface {
	Save(question []Question) ([]Question, error)
	GetAll() ([]Question, error)
	GetAllByQuizId(quizId string) ([]Question, error)
	GetById(id string) (Question, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(question []Question) ([]Question, error) {
	err := r.db.Create(&question).Error

	if err != nil {
		return question, err
	}

	return question, nil
}

func (r *repository) GetAll() ([]Question, error) {
	var questions []Question

	error := r.db.Find(&questions).Error

	if error != nil {
		return questions, error
	}

	return questions, nil
}

func (r *repository) GetAllByQuizId(quizId string) ([]Question, error) {
	var questions []Question

	error := r.db.Where("quiz_id=?", quizId).Find(&questions).Error

	if error != nil {
		return questions, error
	}

	return questions, nil
}

func (r *repository) GetById(id string) (Question, error) {
	var question Question
	err := r.db.Where("id=?", id).Find(&question).Error

	if err != nil {
		return question, err
	}

	return question, nil
}
