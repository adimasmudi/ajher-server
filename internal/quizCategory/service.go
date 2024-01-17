package quizCategory

import (
	"errors"
	"time"
)

type Service interface {
	Save(input QuizCategoryInput) (QuizCategory, error)
	Update(input QuizCategoryInput) (QuizCategory, error)
	Delete(ID string) (QuizCategory, error)
	GetById(ID string) (QuizCategory, error)
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
	quizCategory.CreatedAt = time.Now()
	quizCategory.UpdatedAt = time.Now()

	newQuizCategory, err := s.repository.Save(quizCategory, "quizCategories")

	if err != nil {
		return quizCategory, err
	}

	return newQuizCategory, nil
}

func (s *service) Update(input QuizCategoryInput) (QuizCategory, error) {
	quizCategory := QuizCategory{}

	quizCategory.CategoryName = input.CategoryName
	quizCategory.Description = input.Description
	quizCategory.UpdatedAt = time.Now()

	newQuizCategory, err := s.repository.Update(quizCategory, "quizCategories")

	if err != nil {
		return quizCategory, err
	}

	return newQuizCategory, nil
}

func (s *service) Delete(ID string) (QuizCategory, error) {
	quizCategory, err := s.repository.GetById(ID, "quizCategories")

	if err != nil {
		return quizCategory, errors.New("quiz category with the given id does't exist")
	}

	deletedQuizCategory, err := s.repository.Delete(ID, "quizCategories")

	if err != nil {
		return quizCategory, err
	}

	return deletedQuizCategory, nil
}

func (s *service) GetById(ID string) (QuizCategory, error) {
	quizCategory, err := s.repository.GetById(ID, "quizCategories")

	if err != nil {
		return quizCategory, errors.New("quiz category with the given id does't exist")
	}

	return quizCategory, nil
}

func (s *service) GetAll() ([]QuizCategory, error) {
	quizCategories, err := s.repository.GetAll("quizCategories")

	if err != nil {
		return quizCategories, err
	}

	return quizCategories, nil
}
