package quiz

type CreateQuizInput struct {
	QuizCategoryId string `json:"quiz_category_id" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description" binding:"required"`
}
