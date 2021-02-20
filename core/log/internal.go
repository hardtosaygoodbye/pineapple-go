package log

import (
	"go.uber.org/zap"
)

var HttpLogger, ErrorLogger, InitLogger, GormLogger *zap.Logger

func initInternalLog() {
	GormLogger = getLogger("mysql", "json", "debug", false)
	HttpLogger = getLogger("request", "console", "info", false)
	ErrorLogger = getLogger("panic", "console", "error", false)
	InitLogger = getLogger("init", "console", "info", false)
}

func flushInternalLog() {
	GormLogger.Sync()
	HttpLogger.Sync()
	ErrorLogger.Sync()
	InitLogger.Sync()
}
