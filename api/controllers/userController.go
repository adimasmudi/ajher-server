package controllers

import (
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Signup  godoc
//
// @Summary  signup
// @Description Adding new user to the database
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param   registerUserInput body  user.RegisterUserInput true "User Data"
// @Success  200   {object} user.User
// @Router   user/register [post]
func (h *userHandler) Register(ctx *gin.Context) {
	var input user.RegisterUserInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if input.Password != input.ConfirmPassword {
		errorMessage := gin.H{"errors": "password and confirm password is not same"}
		response := utils.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		response := utils.APIResponse("Register Account Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "abc")

	response := utils.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
