package otp

type VerifyOtpInput struct {
	OtpCode string `json:"otp_code" binding:"required"`
}
