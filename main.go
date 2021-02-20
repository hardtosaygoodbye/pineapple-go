package main

import (
	"pineapple-go/config"
	"pineapple-go/core/database"
	"pineapple-go/core/log"
	"pineapple-go/core/redis"
	"pineapple-go/middleware"
	"pineapple-go/router"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig() //加载配置
}

func main() {
	r := gin.Default()

	log.InitLog()        //配置日志
	redis.ConnectRedis() //连接redis
	database.ConnectDB() //连接数据库
	defer database.DisconnectDB()
	database.AutoMigrate() //自动迁移

	middleware.LoadMiddlewares(r) //加载中间件

	router.LoadRoutes(r)   //加载路由
	r.Run(config.App.Port) // listen and serve on 0.0.0.0:8080
}
