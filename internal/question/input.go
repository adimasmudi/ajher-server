package question

type AddQuestionInputs struct {
	QuizId    string             `json:"quiz_id" binding:"required"`
	Questions []AddQuestionInput `json:"questions" binding:"required"`
}

type AddQuestionInput struct {
	Question        string  `json:"question" binding:"required"`
	ReferenceAnswer string  `json:"reference_answer"`
	GradePercentage float64 `json:"grade_percentage" binding:"required"`
	Duration        string  `json:"duration" binding:"required"`
}
