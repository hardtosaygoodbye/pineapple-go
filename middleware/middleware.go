package middleware

import (
	"pineapple-go/core/log"

	"github.com/gin-gonic/gin"
)

//LoadMiddlewares load all middleware to route
func LoadMiddlewares(router *gin.Engine) {

	router.Use(injectData) //middleware for inject data

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(accessLog(log.HttpLogger))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	// router.Use(recoveryWithLog(log.ErrorLogger, true))
	router.Use(flushLog)

	log.InitLogger.Info("load all middleware successful")
}
