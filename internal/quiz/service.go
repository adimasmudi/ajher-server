package quiz

import (
	"ajher-server/internal/participantQuestion"
	"ajher-server/internal/participation"
	"ajher-server/internal/question"
	"ajher-server/utils"
	"time"
)

type Service interface {
	Save(input CreateQuizInput, userId string) (Quiz, error)
	GetQuizDetail(id string) (Quiz, participation.Participation, error)
	JoinQuiz(quizCode string, userId string) (Quiz, error)
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

func (s *service) Save(input CreateQuizInput, userId string) (Quiz, error) {
	quiz := Quiz{}

	quizCode := utils.GenerateRandomString(8)

	quiz.QuizCategoryId = input.QuizCategoryId
	quiz.Title = input.Title
	quiz.Description = input.Description
	quiz.Status = "active"
	quiz.Code = quizCode
	quiz.CreatedAt = time.Now()

	newQuiz, err := s.repository.Save(quiz, "quizzes")

	if err != nil {
		return quiz, err
	}

	participation := participation.Participation{
		ID:        utils.GenerateRandomString(15),
		UserId:    userId,
		QuizId:    newQuiz.ID,
		Status:    "creator",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.participationRepository.Save(participation, "participations")

	if err != nil {
		return quiz, err
	}

	return newQuiz, nil

}

func (s *service) GetQuizDetail(id string) (Quiz, participation.Participation, error) {
	var participation participation.Participation
	quiz, err := s.repository.GetById(id, "quizzes")

	if err != nil {
		return quiz, participation, err
	}

	participation, err = s.participationRepository.GetByQuizId(quiz.ID, "participations")

	if err != nil {
		return quiz, participation, err
	}

	return quiz, participation, nil
}

func (s *service) JoinQuiz(quizCode string, userId string) (Quiz, error) {
	quiz, err := s.repository.GetByCode(quizCode, "quizzes")

	if err != nil {
		return quiz, err
	}
	questions, err := s.questionRepository.GetAllByQuizId(quiz.ID, "questions")

	if err != nil {
		return quiz, err
	}

	questions = ShuffleArray(questions)

	var newParticipant participation.Participation

	newParticipant.QuizId = quiz.ID
	newParticipant.UserId = userId
	newParticipant.Status = "participant"
	newParticipant.CreatedAt = time.Now()
	newParticipant.UpdatedAt = time.Now()

	savedParticipant, err := s.participationRepository.Save(newParticipant, "participations")

	if err != nil {
		return quiz, err
	}

	var newParticipantQuestions []participantQuestion.ParticipantQuestion

	for idx, item := range questions {
		var newQuestion participantQuestion.ParticipantQuestion

		newQuestion.ParticipationId = savedParticipant.ID
		newQuestion.QuestionId = item.ID
		newQuestion.Number = int64(idx + 1)
		newQuestion.CreatedAt = time.Now()
		newQuestion.UpdatedAt = time.Now()

		newParticipantQuestions = append(newParticipantQuestions, newQuestion)
	}

	_, err = s.participantQuestionRepository.Save(newParticipantQuestions, "participantQuestions")

	if err != nil {
		return quiz, err
	}

	return quiz, nil
}
