package main

import (
	"gindemo/config"
	"gindemo/controllers"
	"gindemo/db"
	"gindemo/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 创建 Gin 引擎
	router := gin.Default()
	config.LoadConfig()
	// 创建一个 loggers 的 Logger 实例
	logger := logrus.New()

	logDir := "log"

	// 检查文件夹是否存在
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		// 创建文件夹
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Fatalf("无法创建日志文件夹：%s", err)
		}
	}
	// 定义日志文件路径
	logPath := filepath.Join("log", "app_request.log")
	// 打开日志文件
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 设置日志输出到文件
		logger.SetOutput(logFile)
	} else {
		logger.Error("无法打开日志文件: ", err)
	}
	db.Initialize()
	// 加载日志中间件

	router.Use(middlewares.LoggerMiddleware(logger))

	// 加载全局异常处理中间件

	router.Use(middlewares.GinGlobalErrorHandler())

	// 实例化控制器
	userController := &controllers.UserController{}
	_ = &controllers.LeadsController{}

	// 创建路由组
	userRoutes := router.Group("/api/v1/business")
	userRoutes.Use(middlewares.GetDBConnection)

	{
		userRoutes.POST("user", userController.CreateUser)
	}

	// 创建路由组
	//orderRoutes := router.Group("/api/v1/business")
	//{
	//orderRoutes.GET("getLeads", orderController.GetOrderList)
	//}

	// 启动服务器
	err = router.Run(":8100")
	if err != nil {
		return
	}
}
