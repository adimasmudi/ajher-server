package quiz

import (
	"ajher-server/internal/participation"
	"ajher-server/utils"
)

type Service interface {
	Save(input CreateQuizInput, userId int) (Quiz, error)
	// GetById(id string) (Quiz, error)
}

type service struct {
	repository              Repository
	participationRepository participation.Repository
}

func NewService(repository Repository, participationRepository participation.Repository) *service {
	return &service{repository, participationRepository}
}

func (s *service) Save(input CreateQuizInput, userId int) (Quiz, error) {
	quiz := Quiz{}

	quizId, err := utils.GeneratedUUID()

	if err != nil {
		return quiz, err
	}

	quizCode := utils.GenerateRandomString(8)

	quiz.ID = quizId
	quiz.QuizCategoryId = input.QuizCategoryId
	quiz.Title = input.Title
	quiz.Description = input.Description
	quiz.Status = "active"
	quiz.Code = quizCode

	newQuiz, err := s.repository.Save(quiz)

	if err != nil {
		return quiz, err
	}

	participation := participation.Participation{
		ID:     utils.GenerateRandomString(15),
		UserId: userId,
		QuizId: newQuiz.ID,
		Status: "creator",
	}

	_, err = s.participationRepository.Save(participation)

	if err != nil {
		return quiz, err
	}

	return newQuiz, nil

}

// func (s *service) GetById(id string) (Quiz, error){

// }
