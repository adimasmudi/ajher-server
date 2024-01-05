package controllers

import (
	"ajher-server/internal/quiz"
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quizHandler struct {
	quizService quiz.Service
}

func NewQuizHandler(quizService quiz.Service) *quizHandler {
	return &quizHandler{quizService}
}

// Save Quiz  godoc
//
// @Summary  save quiz
// @Description Adding new quiz to the database
// @Tags   Quiz
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   createQuizInput body  quiz.CreateQuizInput true "User Data"
// @Success  200   {object} quiz.Quiz
// @Router   /quiz/save [post]
func (h *quizHandler) Save(ctx *gin.Context) {
	var input quiz.CreateQuizInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Save Quiz Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	newQuiz, err := h.quizService.Save(input, userID)

	if err != nil {
		response := utils.APIResponse("Save Quiz Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Save Quiz Category Success", http.StatusOK, "success", newQuiz)

	ctx.JSON(http.StatusOK, response)

}
