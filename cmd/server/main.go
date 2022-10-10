package main

import (
	"github.com/fabiotavarespr/jwt-authentication-golang/internal/controllers"
	"github.com/fabiotavarespr/jwt-authentication-golang/internal/database"
	"github.com/fabiotavarespr/jwt-authentication-golang/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("jwt_demo:golang@tcp(127.0.0.1:3306)/jwt_demo?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
