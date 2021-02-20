package log

import (
	"context"
	"fmt"
	"pineapple-go/config"
	"runtime"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func initAppLog() {
	logger = getLogger("app", config.Log.Formatter, config.Log.LogLevel, true).Sugar()
}

func flushAppLog() {
	logger.Sync()
}

func getData(ctx context.Context, data interface{}) []interface{} {
	slice := make([]interface{}, 0)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("can not get file path and line")
	}

	fileAndLine := fmt.Sprintf("%s:%d", file, line)
	//gorm的file写在data里面，替换到外面标准格式
	if m, ok := data.(map[string]interface{}); ok {
		value, ok := m["file"]
		if ok {
			fileAndLine = value.(string)
			delete(m, "file")

		}
	}
	if ctx.Value("TraceId") != nil {
		slice = append(slice, "TraceId", ctx.Value("TraceId"))
		slice = append(slice, "requestId", ctx.Value("requestId"))
		slice = append(slice, "file", trimmedPath(fileAndLine))
	}
	slice = append(slice, "data", data)
	return slice
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(ctx context.Context, keywords string, value interface{}) {

	logger.Debugw(keywords, getData(ctx, value)...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(ctx context.Context, keywords string, value interface{}) {
	logger.Infow(keywords, getData(ctx, value)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(ctx context.Context, keywords string, value interface{}) {
	logger.Warnw(keywords, getData(ctx, value)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(ctx context.Context, keywords string, value interface{}) {
	logger.Errorw(keywords, getData(ctx, value)...)
}
