package controller

import (
	"pineapple-go/service"

	"github.com/gin-gonic/gin"
)

func QiniuToken(ctx *gin.Context) {
	upToken := service.QiniuService.GetToken()
	Success(ctx, gin.H{
		"uptoken": upToken,
	})
}
