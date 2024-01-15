package answer

import "ajher-server/utils"

type Service interface {
	Save(input AnswerQuestionInput, userId int) (Answer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input AnswerQuestionInput, userId int) (Answer, error) {
	answer := Answer{}

	uuid, err := utils.GeneratedUUID()

	if err != nil {
		return answer, err
	}

	answer.ID = uuid
	answer.UserId = userId
	answer.QuestionId = input.QuestionId
	answer.Answer = input.Answer
	answer.AnswerDuration = input.AnswerDuration
	answer.Status = "process"

	savedAnswer, err := s.repository.Save(answer)

	if err != nil {
		return answer, err
	}

	return savedAnswer, nil
}
