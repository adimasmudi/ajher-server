package answer

type AnswerQuestionInput struct {
	QuestionId     string `json:"question_id"`
	Answer         string `json:"answer"`
	AnswerDuration int    `json:"answer_duration"`
}
