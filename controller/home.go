package controller

import (
	"pineapple-go/core/log"
	"pineapple-go/core/redis"
	"pineapple-go/model"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	log.Debug(ctx, "tt", "this is debug")
	log.Info(ctx, "aaa", "this is ifno")
	log.Warn(ctx, "wa", "this is warn")
	redis.Client.Set(ctx, "name", "Hello JSON", 5*time.Minute).Err()
	log.Error(ctx, "err", "this is error")
	Success(ctx, gin.H{"hello": "test"})
}

func TestRedis(ctx *gin.Context) {
	err := redis.Client.Set(ctx, "name", "Hello JSON", 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	val, err := redis.Client.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	Success(ctx, gin.H{"test": val})
}

// func TestDB(ctx *gin.Context) {
// 	user := model.User{
// 		Username: "范兄弟",
// 		Password: "3333",
// 		CreateAt: time.Now(),
// 	}
// 	err := service.UserService.Create(ctx, &user)
// 	fmt.Print(user)
// 	if err != nil {
// 		log.Error(ctx, "create User error", err)
// 	} else {
// 		log.Info(ctx, "create User", user)
// 		redis.Client.Set(ctx, "hello", user.Username, 5)
// 		Success(ctx, "成功", gin.H{"data": user})
// 	}
// }

func Test(ctx *gin.Context) {
}

func TestQuery(ctx *gin.Context) {

	m := ctx.QueryMap("map")
	value, ok := m["map"]
	if !ok {
		value = "default value"
	}
	Success(ctx, gin.H{
		"name": ctx.Query("name"),
		"age":  ctx.DefaultQuery("age", "default value"),
		"body": ctx.DefaultPostForm("body", "default body"),
		"map":  value,
	})
}

func TestBind(ctx *gin.Context) {
	var user model.User
	//ShouldBindQuery
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(200, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
