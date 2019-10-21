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

func mysqlDBPool() *base.SQLConnPool {
	host := `127.0.0.1:3306`
	database := `wsdk_admin`
	user := `root`
	password := ``
	charset := `utf8`
	// 用于设置最大打开的连接数
	maxOpenConns := 100
	// 用于设置闲置的连接数
	maxIdleConns := 50
	mysqlDB := base.InitMySQLPool(host, database, user, password, charset, maxOpenConns, maxIdleConns)
	return mysqlDB
}
func main() {
	mysqlDB := mysqlDBPool()
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

	router := gin.Default()
	api := router.Group("/v1")

	user := api.Group("user")
	{
		user.GET("/", func(context *gin.Context) {
			controllers.GetUserController.Index(context, mysqlDB)
		})
		user.GET("/create/:username", func(context *gin.Context) {
			controllers.GetUserController.Create(context, mysqlDB)
		})
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
