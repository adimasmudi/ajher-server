package routes

import (
	"ajher-server/api/controllers"
	"ajher-server/api/middleware"
	"ajher-server/internal/otp"
	"ajher-server/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(api *gin.RouterGroup, database *gorm.DB) {
	userRoute := api.Group("user")

	// repository
	userRepository := user.NewRepository(database)
	otpRepository := otp.NewRepository(database)

	// services
	userService := user.NewService(userRepository, otpRepository)
	otpService := otp.NewService(otpRepository)

	// controllers
	userHandler := controllers.NewUserHandler(userService, otpService)

	// auth middleware
	authMiddleware := middleware.NewAuthMiddleware(userService)

	userRoute.POST("/register", userHandler.Register)
	userRoute.POST("/login", userHandler.Login)
	userRoute.GET("/profile", authMiddleware.AuthMiddleware, userHandler.GetProfile)
	userRoute.GET("/validateToken", authMiddleware.AuthMiddleware, userHandler.ValidateToken)
	userRoute.POST("/refreshToken", authMiddleware.RefreshTokenMiddleware, userHandler.RefreshToken)
	userRoute.POST("/googleAuth", userHandler.GoogleAuth)
	userRoute.POST("/resetPassword", userHandler.ResetPassword)
	userRoute.POST("/verifyOtp", userHandler.VerifyOtp)
	userRoute.POST("/changePassword", userHandler.ChangePassword)
}
