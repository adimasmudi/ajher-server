package controllers

import (
	"ajher-server/internal/quizCategory"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quizCategoryHandler struct {
	quizCategoryService quizCategory.Service
}

func NewQuizCategoryHandler(quizCategoryService quizCategory.Service) *quizCategoryHandler {
	return &quizCategoryHandler{quizCategoryService}
}

// Save Quiz Category  godoc
//
// @Summary  save quiz category
// @Description Adding new quiz category to the database
// @Tags   Quiz Category
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   quizCategoryInput body  quizCategory.QuizCategoryInput true "User Data"
// @Success  200   {object} quizCategory.QuizCategory
// @Router   /quizCategory/save [post]
func (h *quizCategoryHandler) Save(ctx *gin.Context) {
	var input quizCategory.QuizCategoryInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Save Quiz Category Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newQuizCategory, err := h.quizCategoryService.Save(input)

	if err != nil {
		response := utils.APIResponse("Save Quiz Category Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Save Quiz Category Success", http.StatusOK, "success", newQuizCategory)

	ctx.JSON(http.StatusOK, response)
}

// get all quiz category  godoc
//
// @Summary  get all quiz category
// @Description Get all quiz category
// @Tags   Quiz Category
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success  200   {object} quizCategory.QuizCategory
// @Router   /quizCategory [get]
func (h *quizCategoryHandler) GetAll(ctx *gin.Context) {

	quizCategories, err := h.quizCategoryService.GetAll()

	if err != nil {
		response := utils.APIResponse("Get Quiz Category Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Get Quiz Category Success", http.StatusOK, "success", quizCategories)
	ctx.JSON(http.StatusOK, response)
}

// get quiz category by id  godoc
//
// @Summary  get quiz category by id
// @Description Get quiz category by id
// @Tags   Quiz Category
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "Quiz Category Id"
// @Success  200   {object} quizCategory.QuizCategory
// @Router   /quizCategory/{id} [get]
func (h *quizCategoryHandler) GetById(ctx *gin.Context) {
	quizCategoryId := ctx.Param("id")

	quizCategory, err := h.quizCategoryService.GetById(quizCategoryId)

	if err != nil {
		response := utils.APIResponse("Get Quiz Category Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Get Quiz Category Success", http.StatusOK, "success", quizCategory)
	ctx.JSON(http.StatusOK, response)
}
