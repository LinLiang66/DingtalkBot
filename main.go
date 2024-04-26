package main

import (
	"DingtalkBot/config"
	"DingtalkBot/handlers"
	"DingtalkBot/routers"
	"DingtalkBot/utils"
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"time"
)

//go:embed config/app.json
var configjson string

func main() {
	conf, err := config.ChangeConfig(configjson)
	if err != nil {
		log.Printf("读取配置文件失败: %v", err.Error())
	}
	utils.InitRedisUtil(conf.RedisConfig.Addr, conf.RedisConfig.Port, conf.RedisConfig.Password)
	handlers.InitHandlers()
	logger := enableLog()
	defer utils.CloseLogger(logger)
	// 注册处理器 默认开启日志打印
	//g := gin.Default()
	// 注册处理器 默认关闭日志打印
	g := gin.New()
	//设置日志级别为 gin.DebugLevelNone，不打印请求路径日志
	g.Use(utils.CustomMiddleware())

	g.GET("/ping", func(c *gin.Context) {

		c.Header("Server", "Go-Gin-Server")
		c.JSON(200, gin.H{
			"message":   "pong",
			"code":      200,
			"success":   true,
			"timestamp": time.Now().UnixNano() / int64(time.Millisecond),
		})
	})

	//添加路由
	routers.RegisterRouter(g)
	// 启动服务
	err = g.Run(":" + conf.AppPort) //
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func enableLog() *lumberjack.Logger {
	// Set up the logger
	var logger *lumberjack.Logger
	logger = &lumberjack.Logger{
		Filename: "logs/go_robot.log",
		MaxSize:  100,      // megabytes
		MaxAge:   365 * 10, // days
	}

	fmt.Printf("logger %T\n", logger)

	// Set up the logger to write to both file and console
	log.SetOutput(io.MultiWriter(logger, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
	// Write some log messages
	log.Println("Starting application...")

	return logger
}
