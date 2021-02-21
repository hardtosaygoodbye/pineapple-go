package router

import (
	"fmt"
	"pineapple-go/controller"
	"pineapple-go/middleware"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) {

	// gin.DisableConsoleColor()
	//404错误
	router.NoRoute(controller.NoRoute)

	router.POST("/wxlogin", controller.WxLogin)
	router.POST("/phone_login", controller.PhoneLogin)
	router.POST("/auth_code", controller.AuthCode)

	router.GET("/client_connect", controller.ClientConnect)
	router.POST("/v1/sues/course", controller.SuesCourse)
	v1 := router.Group("/v1")
	v1.Use(middleware.CheckHeaderRequired, middleware.JWTAuthRequired)
	{
		// v1.GET("/ping", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "pong",
		// 	})
		// })
		v1.POST("/user/user_info", controller.UserInfo)
		v1.POST("/user/update_user_info", controller.UpdateUserInfo)
		v1.POST("/weico/publish", controller.PublishWeico)
		v1.POST("/weico/delete", controller.DeleteWeico)
		v1.POST("/weico/list", controller.WeicoList)
		v1.POST("/weico/like", controller.LikeWeico)
		v1.POST("/weico/add_comment", controller.CommentWeico)
		v1.POST("/weico/comment_list", controller.CommentList)
		v1.POST("/weico/delete_comment", controller.DeleteComment)
		v1.POST("/weico/cate_list", controller.WeicoCateList)
		v1.POST("/qiniu/token", controller.QiniuToken)
		// v1.GET("/home", controller.Index)
		// v1.GET("/testredis", controller.TestRedis)
		// // v1.GET("/testdb", controller.TestDB)
		// v1.GET("/test", controller.Test)
		// v1.GET("/query", controller.TestQuery)
		// v1.GET("/bind", controller.TestBind)
		// v1.GET("/userinfo", controller.UserInfo)
	}

	fmt.Println("load routes successful.")
}
