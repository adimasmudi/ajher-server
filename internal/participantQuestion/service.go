package participantQuestion

import (
	"ajher-server/internal/participation"
	"ajher-server/internal/user"
	"errors"
)

type Service interface {
	GetQuestionByEachNumber(userId int, quizId string) ([]ParticipantQuestion, error)
}

type service struct {
	repository              Repository
	userRepository          user.Repository
	participationRepository participation.Repository
}

func NewService(repository Repository, userRepository user.Repository, participationRepository participation.Repository) *service {
	return &service{repository, userRepository, participationRepository}
}

func (s *service) GetQuestionByEachNumber(userId int, quizId string) ([]ParticipantQuestion, error) {
	var participationQuestion []ParticipantQuestion
	user, err := s.userRepository.GetById(userId)

	if err != nil {
		return participationQuestion, errors.New("user doesn't exist")
	}

	participation, err := s.participationRepository.GetByUserId(user.ID)

	if err != nil {
		return participationQuestion, errors.New("this user never participated yet")
	}

	participationQuestion, err = s.repository.GetByParticipantId(participation.ID)

	if err != nil {
		return participationQuestion, err
	}

	for _, data := range participationQuestion {
		if data.Question.QuizId != quizId {
			return participationQuestion, errors.New("question with the given quiz id is not available")
		}
	}

	return participationQuestion, nil
}
