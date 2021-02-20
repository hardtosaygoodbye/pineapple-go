package middleware

import (
	"net"
	"os"
	"pineapple-go/controller"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func recoveryWithLog(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					logger.Error("[panic error]",
						zap.Any("error", err),
						zap.String("url", c.Request.URL.Path),
						zap.String("request", c.Request.URL.RawQuery),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[panic error]",
						zap.Any("error", err),
						zap.String("url", c.Request.URL.Path),
						zap.String("request", c.Request.URL.RawQuery),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[panic error]",
						zap.Any("error", err),
						zap.String("url", c.Request.URL.Path),
						zap.String("request", c.Request.URL.RawQuery),
					)
				}
				controller.StatusInternalServerError(c)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
