package main

import (
	"base"
	"controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)
	// Logging to a file.
	f, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//加载配置文件
	if base.INIT_OBJ.Init() {
		log.Println("====== 配置文件加载成功 ======")
	}
	//初始化数据库
	//if base.MODEL.InitDB() {
	//	log.Println("====== 数据库初始化成功 ======")
	//}

	router := gin.Default()
	api := router.Group("/v1")

	user := api.Group("user")
	{
		user.GET("/", controllers.GetUserController.Index)
		user.GET("/create/:username", controllers.GetUserController.Create)
	}

	about := api.Group("about")
	{
		about.GET("/", controllers.GetAboutController.Index)
	}

	err := router.Run("0.0.0.0:8081")
	if err != nil {
		fmt.Println("启动错误")
	}
}
