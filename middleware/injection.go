package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func injectData(ctx *gin.Context) {
	requestId := uuid.New().String()
	ctx.Set("requestId", requestId)
}
