package controllers

import (
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Token Object
type TokenObject struct {
	AccessToken  string
	RefreshToken string
}

// Register  godoc
//
// @Summary  register
// @Description Adding new user to the database
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param   registerUserInput body  user.RegisterUserInput true "User Data"
// @Success  200   {object} TokenObject
// @Router   /user/register [post]
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

	token, err := utils.GenerateToken(newUser.ID)

	if err != nil {
		response := utils.APIResponse("Register Account Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Account has been registered", http.StatusOK, "success", token)

	ctx.JSON(http.StatusOK, response)
}

// login  godoc
//
// @Summary  login
// @Description Login user
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param   loginUserInput body  user.LoginUserInput true "User Data"
// @Success  200   {object} TokenObject
// @Router   /user/login [post]
func (h *userHandler) Login(ctx *gin.Context) {
	var input user.LoginUserInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logedinUser, err := h.userService.Login(input)

	if err != nil {
		response := utils.APIResponse("Login Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := utils.GenerateToken(logedinUser.ID)

	if err != nil {
		response := utils.APIResponse("Login Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Login success", http.StatusOK, "success", token)

	ctx.JSON(http.StatusOK, response)
}

// google auth  godoc
//
// @Summary  google auth
// @Description Google Auth
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param   googleOAuthInput body  user.GoogleOAuthInput true "User Data"
// @Success  200   {object} TokenObject
// @Router   /user/googleAuth [post]
func (h *userHandler) GoogleAuth(ctx *gin.Context) {
	var input user.GoogleOAuthInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Google Auth Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logedinUser, err := h.userService.GoogleAuth(input)

	if err != nil {
		response := utils.APIResponse("Google Auth Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := utils.GenerateToken(logedinUser.ID)

	if err != nil {
		response := utils.APIResponse("Google Auth Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Google Auth success", http.StatusOK, "success", token)

	ctx.JSON(http.StatusOK, response)
}

// get profile  godoc
//
// @Summary  get profile
// @Description Get user profile
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success  200   {object} user.User
// @Router   /user/profile [get]
func (h *userHandler) GetProfile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	userProfile, err := h.userService.GetUserById(userID)

	if err != nil {
		response := utils.APIResponse("Get Profile Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userProfile)

	response := utils.APIResponse("Get Profile success", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

// validate token  godoc
//
// @Summary  validate token
// @Description Validate JWT token
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success  200   {object} TokenObject
// @Router   /user/validateToken [get]
func (h *userHandler) ValidateToken(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	user, err := h.userService.GetUserById(userID)

	if err != nil {
		response := utils.APIResponse("Token is invalid", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		response := utils.APIResponse("Validate token failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Validate token success", http.StatusOK, "success", token)

	ctx.JSON(http.StatusOK, response)
}

// refresh token  godoc
//
// @Summary  refresh token
// @Description Refresh JWT token
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   refreshTokenInput body  user.RefreshTokenInput true "User Data"
// @Success  200   {object} TokenObject
// @Router   /user/refreshToken [post]
func (h *userHandler) RefreshToken(ctx *gin.Context) {
	var input user.RefreshTokenInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Refresh Token Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := utils.ValidateToken(input.AccessToken)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Refresh Token Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		response := utils.APIResponse("Refresh Token Faild", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userId := int(claims["user_id"].(float64))

	user, err := h.userService.GetUserById(userId)

	if err != nil {
		response := utils.APIResponse("Refresh Token Failed", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	newToken, err := utils.GenerateToken(user.ID)

	if err != nil {
		response := utils.APIResponse("Refresh Token Failed", http.StatusUnauthorized, "error", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	ctx.Set("currentUser", user)

	response := utils.APIResponse("Refresh Token success", http.StatusOK, "success", newToken)

	ctx.JSON(http.StatusOK, response)
}

// Reset Password  godoc
//
// @Summary  reset password
// @Description Send email consist of OTP to user
// @Tags   User Authentication
// @Accept   json
// @Produce  json
// @Param   resetPasswordInput body  user.ResetPasswordInput true "User Data"
// @Success  200   {object} TokenObject
// @Router   /user/resetPassword [post]
func (h *userHandler) ResetPassword(ctx *gin.Context) {
	var input user.ResetPasswordInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Send Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	otp, err := h.userService.GenerateAndSendEmail(input)

	if err != nil {
		response := utils.APIResponse("Send Email Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Send Email Success", http.StatusOK, "success", otp)

	ctx.JSON(http.StatusOK, response)
}
