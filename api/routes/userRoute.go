package routes

import (
	"ajher-server/api/controllers"
	"ajher-server/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(api *gin.RouterGroup, database *gorm.DB) {
	userRoute := api.Group("user")

	// repository
	userRepository := user.NewRepository(database)

	// services
	userService := user.NewService(userRepository)

	// controllers
	userHandler := controllers.NewUserHandler(userService)

	userRoute.POST("/register", userHandler.Register)
}
