package main

import (
	"log"

	"ajher-server/database"
	"ajher-server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title  Ajher API
// @version  1.0
// @description API for ajher apps

// @securityDefinitions.apiKey JWT
// @in       header
// @name      token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /
// @schemes http
func main() {
	docs.SwaggerInfo.Title = "Ajher API"
	docs.SwaggerInfo.Description = "API for ajher apps"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":5000")

}
