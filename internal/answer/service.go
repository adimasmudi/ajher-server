package answer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Service interface {
	Save(input AnswerQuestionInput, userId string) (Answer, error)
	FinishAnswer(quizId string, userId string) ([]Answer, error)
	GetFinishedAnswer(quizId string, userID string) ([]AnswerWithQuestion, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input AnswerQuestionInput, userId string) (Answer, error) {
	answer := Answer{}

	answer.UserID = userId
	answer.QuestionID = input.QuestionId
	answer.AnswerText = input.Answer
	answer.AnswerDuration = input.AnswerDuration
	answer.Status = "process"
	answer.CreatedAt = time.Now()
	answer.UpdatedAt = time.Now()

	savedAnswer, err := s.repository.Save(answer, "answers")

	if err != nil {
		return answer, err
	}

	return savedAnswer, nil
}

func (s *service) FinishAnswer(quizId string, userId string) ([]Answer, error) {
	var answerToReturn []Answer
	type answerCorrectionRequest struct {
		Answers []AnswerToCorrect `json:"answers"`
	}
	var answerToCorrect AnswerToCorrect
	var answerToCorrects []AnswerToCorrect

	userAnswers, err := s.repository.GetUserAnswers(quizId, userId)

	if err != nil {
		return nil, err
	}

	for _, data := range userAnswers {
		answerToCorrect.AnswerId = data.ID
		answerToCorrect.Answer = data.AnswerText
		answerToCorrect.QuestionId = data.Question.ID
		answerToCorrect.ReferenceAnswer = data.Question.ReferenceAnswer
		answerToCorrect.AnswerDuration = data.Answer.AnswerDuration

		answerToCorrects = append(answerToCorrects, answerToCorrect)
	}

	answerRequest := answerCorrectionRequest{
		Answers: answerToCorrects,
	}

	modelUrl := os.Getenv("MODEL_SERVICE_URL")
	formData, err := json.Marshal(answerRequest)

	if err != nil {
		return answerToReturn, err
	}

	reader := bytes.NewReader(formData)

	req, err := http.NewRequest("POST", modelUrl, reader)

	if err != nil {
		return answerToReturn, err
	}

	// Set headers (adjust content type as needed)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return answerToReturn, err
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return answerToReturn, err
	}

	type responseStruct struct {
		QuestionId      string  `json:"question_id"`
		AnswerId        string  `json:"answer_id"`
		Answer          string  `json:"answer"`
		ReferenceAnswer string  `json:"reference_answer"`
		Grade           float64 `json:"grade"`
		AnswerDuration  int64   `json:"answer_duration"`
	}

	type responseObject struct {
		Message string           `json:"message"`
		Code    int64            `json:"code"`
		Status  string           `json:"status"`
		Data    []responseStruct `json:"data"`
	}

	var responseData responseObject
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return answerToReturn, err
	}

	var answerPayloads []Answer
	var answerPayload Answer

	for _, data := range responseData.Data {
		answerPayload.QuestionID = data.QuestionId
		answerPayload.ID = data.AnswerId
		answerPayload.UserID = userId
		answerPayload.Grade = data.Grade
		answerPayload.AnswerText = data.Answer
		answerPayload.AnswerDuration = data.AnswerDuration

		var label string
		if data.Grade <= 50 {
			label = "wrong"
		} else {
			label = "right"
		}

		answerPayload.Label = label
		answerPayload.Status = "evaluated"
		answerPayload.UpdatedAt = time.Now()

		answerPayloads = append(answerPayloads, answerPayload)
	}

	updatedAnswer, err := s.repository.Update(answerPayloads, "answers")

	if err != nil {
		return updatedAnswer, err
	}

	return updatedAnswer, nil
}

func (s *service) GetFinishedAnswer(quizId string, userID string) ([]AnswerWithQuestion, error) {
	userAnswers, err := s.repository.GetUserAnswers(quizId, userID)

	if err != nil {
		return nil, err
	}

	return userAnswers, nil
}
