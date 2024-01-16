package controllers

import (
	"ajher-server/internal/answer"
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type answerHandler struct {
	answerService answer.Service
}

func NewAnswerHandler(answerService answer.Service) *answerHandler {
	return &answerHandler{answerService}
}

// Save Answer  godoc
//
// @Summary  save answer
// @Description Adding new answer to the database.
// @Tags   Answer
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   AnswerQuestionInput body  answer.AnswerQuestionInput true "Answer Data"
// @Success  200   {object} answer.Answer
// @Router   /answer/save [post]
func (h *answerHandler) Save(ctx *gin.Context) {
	var input answer.AnswerQuestionInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Save Answer Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	newAnswer, err := h.answerService.Save(input, userID)

	if err != nil {
		response := utils.APIResponse("Save Answer Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Save Answer Success", http.StatusOK, "success", newAnswer)

	ctx.JSON(http.StatusOK, response)
}

// Finish Answer Quiz  godoc
//
// @Summary  finish answer quiz
// @Description Finish answering question in quiz.
// @Tags   Answer
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param quizId path string true "Quiz Id"
// @Success  200   {object} answer.Answer
// @Router   /answer/finish/{quizId} [post]
func (h *answerHandler) FinishAnswer(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	quizId := ctx.Param("quizId")

	finishedQuiz, err := h.answerService.FinishAnswer(quizId, userID)

	if err != nil {
		response := utils.APIResponse("Finish Quiz Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Finish Quiz Success", http.StatusOK, "success", finishedQuiz)

	ctx.JSON(http.StatusOK, response)
}

// get finished answer godoc
//
// @Summary  get finished answer
// @Description Get finished answer
// @Tags  Answer
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param quizId path string true "Quiz Id"
// @Success  200   {object} answer.Answer
// @Router   /answer/getFinished/{quizId} [get]
func (h *answerHandler) GetFinished(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	quizId := ctx.Param("quizId")

	finishedAnswer, err := h.answerService.GetFinishedAnswer(quizId, userID)

	if err != nil {
		response := utils.APIResponse("Get Finished Answer Quiz Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := answer.FormatFinishAnswer(finishedAnswer)
	response := utils.APIResponse("Get Finished Answer Quiz Success", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
