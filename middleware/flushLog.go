package middleware

import (
	"pineapple-go/core/log"

	"github.com/gin-gonic/gin"
)

func flushLog(ctx *gin.Context) {
	log.LogSync()
}
