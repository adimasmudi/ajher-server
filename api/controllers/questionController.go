package controllers

import (
	"ajher-server/internal/participantQuestion"
	"ajher-server/internal/question"
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService            question.Service
	participantQuestionService participantQuestion.Service
}

func NewQuestionHandler(questionService question.Service, participantQuestionService participantQuestion.Service) *questionHandler {
	return &questionHandler{questionService, participantQuestionService}
}

// Save Questions  godoc
//
// @Summary  save questions
// @Description Adding new questions to the database. Add field duration input as string for example "50 sec" or "1 min", it will converted into second in server and return to client as second also. The client should format it.
// @Tags   Question
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   addQuestionInputs body  question.AddQuestionInputs true "User Data"
// @Success  200   {object} question.Question
// @Router   /question/save [post]
func (h *questionHandler) Save(ctx *gin.Context) {
	var input question.AddQuestionInputs

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Save Questions Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newQuestions, err := h.questionService.Save(input)

	if err != nil {
		response := utils.APIResponse("Save Questions Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Save Questions Success", http.StatusOK, "success", newQuestions)

	ctx.JSON(http.StatusOK, response)
}

// Get Question  godoc
//
// @Summary  get question each number
// @Description Get question by each number from the database
// @Tags   Question
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param quizId path string true "Quiz Id"
// @Success  200   {object} participantQuestion.ParticipantQuestion
// @Router   /question/{quizId} [get]
func (h *questionHandler) GetQuestionByNumber(ctx *gin.Context) {
	quizId := ctx.Param("quizId")

	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	participationQuestion, err := h.participantQuestionService.GetQuestionByEachNumber(userID, quizId)

	if err != nil {
		response := utils.APIResponse("Get Question Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := participantQuestion.FormatQuestion(participationQuestion)
	response := utils.APIResponse("Get Question Success", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
