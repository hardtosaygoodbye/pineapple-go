package middleware

import (
	"pineapple-go/config"
	"pineapple-go/constant"
	"pineapple-go/controller"
	"pineapple-go/core/log"
	"pineapple-go/core/redis"
	"pineapple-go/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func JWTAuthRequired(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authentication")
	uid, err := util.ParseToken(authToken, config.App.JWT_TOKEN)
	if err != nil {
		controller.Error(ctx, constant.USER_JWT_PARSE_FAILD, "jwt 解析失败")
		log.Error(ctx, "auth parse jwt", err)
		ctx.Abort()
		return
	}
	//校验redis中是否存在
	val, err := redis.Client.Exists(ctx, "jwt:user:"+uid).Result()
	if err != nil {
		log.Error(ctx, "auth redis", err)
		controller.Error(ctx, constant.REDIS_KEY_NOT_EXISTS_ERR, "redis 用户不存在")
		ctx.Abort()
		return
	}
	if val <= 0 {
		controller.Error(ctx, constant.REDIS_KEY_NOT_EXISTS_ERR, "redis 用户不存在")
		ctx.Abort()
		return
	}
	tmp, err := strconv.Atoi(uid)
	ctx.Set("uid", tmp)
}
