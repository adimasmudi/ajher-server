package user

type RegisterUserInput struct {
	Email           string `json:"email" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenInput struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

type GoogleOAuthInput struct {
	OAuthAccessToken string `json:"oAuthAccessToken" binding:"required"`
}

type ResetPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ChangePasswordUserInput struct {
	OtpCode         string `json:"otp_code" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"change_password" binding:"required"`
}
