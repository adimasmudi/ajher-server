package quiz

type CreateQuizInput struct {
	QuizCategoryId int    `json:"quiz_category_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}
