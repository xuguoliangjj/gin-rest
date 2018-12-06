package main

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	api := router.Group("/v1")

	user := api.Group("user")
	{
		user.GET("/", controllers.GetUserController.Index)
		user.GET("/create", controllers.GetUserController.Create)
	}

	about := api.Group("about")
	{
		about.GET("/", controllers.GetAboutController.Index)
	}

	router.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8081
}
