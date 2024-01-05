package quizCategory

import (
	"errors"
)

type Service interface {
	Save(input QuizCategoryInput) (QuizCategory, error)
	Update(input QuizCategoryInput) (QuizCategory, error)
	Delete(ID int) (QuizCategory, error)
	GetById(ID int) (QuizCategory, error)
	GetAll() ([]QuizCategory, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input QuizCategoryInput) (QuizCategory, error) {
	quizCategory := QuizCategory{}

	quizCategory.CategoryName = input.CategoryName
	quizCategory.Description = input.Description

	newQuizCategory, err := s.repository.Save(quizCategory)

	if err != nil {
		return quizCategory, err
	}

	return newQuizCategory, nil
}

func (s *service) Update(input QuizCategoryInput) (QuizCategory, error) {
	quizCategory := QuizCategory{}

	quizCategory.CategoryName = input.CategoryName
	quizCategory.Description = input.Description

	newQuizCategory, err := s.repository.Update(quizCategory)

	if err != nil {
		return quizCategory, err
	}

	return newQuizCategory, nil
}

func (s *service) Delete(ID int) (QuizCategory, error) {
	quizCategory, err := s.repository.GetById(ID)

	if err != nil {
		return quizCategory, errors.New("quiz category with the given id does't exist")
	}

	deletedQuizCategory, err := s.repository.Delete(ID)

	if err != nil {
		return quizCategory, err
	}

	return deletedQuizCategory, nil
}

func (s *service) GetById(ID int) (QuizCategory, error) {
	quizCategory, err := s.repository.GetById(ID)

	if err != nil {
		return quizCategory, errors.New("quiz category with the given id does't exist")
	}

	return quizCategory, nil
}

func (s *service) GetAll() ([]QuizCategory, error) {
	quizCategories, err := s.repository.GetAll()

	if err != nil {
		return quizCategories, err
	}

	return quizCategories, nil
}
