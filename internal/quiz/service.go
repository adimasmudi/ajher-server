package quiz

import (
	"ajher-server/internal/participantQuestion"
	"ajher-server/internal/participation"
	"ajher-server/internal/question"
	"ajher-server/utils"
	"fmt"
)

type Service interface {
	Save(input CreateQuizInput, userId int) (Quiz, error)
	GetQuizDetail(id string) (Quiz, participation.Participation, error)
	JoinQuiz(quizCode string, userId int) (Quiz, error)
}

type service struct {
	repository                    Repository
	participationRepository       participation.Repository
	questionRepository            question.Repository
	participantQuestionRepository participantQuestion.Repository
}

func NewService(repository Repository, participationRepository participation.Repository, questionRepository question.Repository, participantQuestionRepository participantQuestion.Repository) *service {
	return &service{repository, participationRepository, questionRepository, participantQuestionRepository}
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

func (s *service) GetQuizDetail(id string) (Quiz, participation.Participation, error) {
	var participation participation.Participation
	quiz, err := s.repository.GetById(id)

	if err != nil {
		return quiz, participation, err
	}

	participation, err = s.participationRepository.GetByQuizId(quiz.ID)

	if err != nil {
		return quiz, participation, err
	}

	return quiz, participation, nil
}

func (s *service) JoinQuiz(quizCode string, userId int) (Quiz, error) {
	quiz, err := s.repository.GetByCode(quizCode)

	if err != nil {
		return quiz, err
	}
	questions, err := s.questionRepository.GetAllByQuizId(quiz.ID)

	if err != nil {
		return quiz, err
	}

	fmt.Println("questions", questions)

	questions = ShuffleArray(questions)

	fmt.Println(questions)

	var newParticipant participation.Participation
	participantId, err := utils.GeneratedUUID()

	if err != nil {
		return quiz, err
	}

	newParticipant.ID = participantId
	newParticipant.QuizId = quiz.ID
	newParticipant.UserId = userId
	newParticipant.Status = "entrant"

	savedParticipant, err := s.participationRepository.Save(newParticipant)

	if err != nil {
		return quiz, err
	}

	var newParticipantQuestions []participantQuestion.ParticipantQuestion

	for idx, item := range questions {
		var newQuestion participantQuestion.ParticipantQuestion

		id, err := utils.GeneratedUUID()

		if err != nil {
			return quiz, err
		}

		newQuestion.ID = id
		newQuestion.ParticipationId = savedParticipant.ID
		newQuestion.QuestionId = item.ID
		newQuestion.Number = idx + 1

		newParticipantQuestions = append(newParticipantQuestions, newQuestion)
	}

	_, err = s.participantQuestionRepository.Save(newParticipantQuestions)

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}
