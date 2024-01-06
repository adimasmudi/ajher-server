package question

import (
	"ajher-server/utils"
)

type Service interface {
	Save(input AddQuestionInputs) ([]Question, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input AddQuestionInputs) ([]Question, error) {
	var questions []Question

	questionsAmount := len(input.Questions)

	for _, questionData := range input.Questions {
		question := Question{}

		questionId, err := utils.GeneratedUUID()

		if err != nil {
			return questions, err
		}

		theDuration, err := utils.ConvertStringDurationIntoInteger(questionData.Duration)

		if err != nil {
			return questions, err
		}

		question.ID = questionId
		question.Question = questionData.Question
		question.Duration = theDuration
		question.GradePercentage = questionData.GradePercentage
		question.ReferenceAnswer = questionData.ReferenceAnswer
		question.Status = "active"
		question.QuizId = input.QuizId

		questionPoint := utils.CalculatePoint(question.Question)

		question.Point = float64(questionsAmount) * float64(questionPoint) * 10

		questions = append(questions, question)
	}

	newQuestions, err := s.repository.Save(questions)

	if err != nil {
		return newQuestions, err
	}

	return newQuestions, nil
}
