package log

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

//DBLogger database logger struct
type DBLogger struct {
	SlowThreshold time.Duration
	gorm_logger.Interface
}

// LogMode log mode
func (l *DBLogger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	newlogger := *l
	return &newlogger
}

// Info print info
func (l DBLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Info(ctx, "gorm info", message)
	} else {
		GormLogger.Info("gorm info", zap.Any("data", message))
	}
}

// Warn print warn messages
func (l DBLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Warn(ctx, "gorm warn", message)
	} else {
		GormLogger.Warn("gorm warn", zap.Any("data", message))
	}

}

// Error print error messages
func (l DBLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Error(ctx, "gorm error", message)
	} else {
		GormLogger.Error("gorm error", zap.Any("data", message))
	}
	panic(message)
}

// Trace print sql message
func (l DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()
	m := map[string]interface{}{
		"file": utils.FileWithLineNum(),
		"cost": float64(elapsed.Nanoseconds()) / 1e9,
		"sql":  sql,
		"rows": rows,
	}

	switch {
	case err != nil:
		m["err"] = err

		if ctx.Value("traceId") != nil {
			Error(ctx, "gorm error", m)
		} else {
			GormLogger.Error("gorm error", zap.Any("data", m))
		}
		panic(m)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		if ctx.Value("traceId") != nil {
			Warn(ctx, "gorm warn", m)
		} else {
			GormLogger.Warn("gorm", zap.Any("data", m))
		}
	default:
		if ctx.Value("traceId") != nil {
			Debug(ctx, "gorm debug", m)
		} else {
			GormLogger.Debug("gorm", zap.Any("data", m))
		}
	}
}
