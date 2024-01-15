package answer

import (
	"ajher-server/utils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Service interface {
	Save(input AnswerQuestionInput, userId int) (Answer, error)
	FinishAnswer(quizId string, userId int) ([]Answer, error)
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

func (s *service) FinishAnswer(quizId string, userId int) ([]Answer, error) {
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
		answerToCorrect.Answer = data.Answer
		answerToCorrect.QuestionId = data.QuestionId
		answerToCorrect.ReferenceAnswer = data.Question.ReferenceAnswer

		answerToCorrects = append(answerToCorrects, answerToCorrect)
	}

	answerRequest := answerCorrectionRequest{
		Answers: answerToCorrects,
	}

	modelUrl := os.Getenv("MODEL_SERVICE_URL")
	formData, err := json.Marshal(answerRequest)

	if err != nil {
		return userAnswers, err
	}

	reader := bytes.NewReader(formData)

	req, err := http.NewRequest("POST", modelUrl, reader)

	if err != nil {
		return userAnswers, err
	}

	// Set headers (adjust content type as needed)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return userAnswers, err
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return userAnswers, err
	}

	type responseStruct struct {
		QuestionId      string  `json:"question_id"`
		AnswerId        string  `json:"answer_id"`
		Answer          string  `json:"answer"`
		ReferenceAnswer string  `json:"reference_answer"`
		Grade           float64 `json:"grade"`
	}

	type responseObject struct {
		Message string           `json:"message"`
		Code    int              `json:"code"`
		Status  string           `json:"status"`
		Data    []responseStruct `json:"data"`
	}

	var responseData responseObject
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return userAnswers, err
	}

	var answerPayloads []Answer
	var answerPayload Answer

	for _, data := range responseData.Data {
		answerPayload.QuestionId = data.QuestionId
		answerPayload.ID = data.AnswerId
		answerPayload.UserId = userId
		answerPayload.Grade = data.Grade

		answerPayloads = append(answerPayloads, answerPayload)
	}

	updatedAnswer, err := s.repository.Update(answerPayloads)

	if err != nil {
		return updatedAnswer, err
	}

	return updatedAnswer, nil
}
