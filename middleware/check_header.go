package middleware

import (
	"pineapple-go/constant"
	"pineapple-go/controller"

	"github.com/gin-gonic/gin"
)

type HeaderField struct {
	Authentication string `json:"Authentication" binding:"required"`
	// TraceId        string `json:"TraceId" binding:"required"`
}

func CheckHeaderRequired(ctx *gin.Context) {
	var h HeaderField
	if err := ctx.ShouldBindHeader(&h); err != nil {
		controller.Error(ctx, constant.CODE_403, err.Error())
		ctx.Abort()
	}
	// ctx.Set("TraceId", h.TraceId)
}
