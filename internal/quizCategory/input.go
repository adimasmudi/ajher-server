package quizCategory

type QuizCategoryInput struct {
	CategoryName string `json:"category_name" binding:"required"`
	Description  string `json:"description" binding:"required"`
}
