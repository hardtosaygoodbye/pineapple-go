package log

import (
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs" //日志切分
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogLevel(level string) zapcore.Level {

	m := map[string]zapcore.Level{
		"info":  zapcore.InfoLevel,
		"debug": zapcore.DebugLevel,
		"error": zapcore.ErrorLevel,
		"warn":  zapcore.WarnLevel,
	}

	l, ok := m[level]
	if ok {
		return l
	} else {
		return zapcore.InfoLevel
	}
}

func InitLog() {
	initAppLog()
	initInternalLog()
	InitLogger.Info("init log successful.")
}

func getLogger(filename string, encodertype string, level string, bapp bool) *zap.Logger {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "file",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "keywords",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder, //zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var encoder zapcore.Encoder
	if encodertype == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var core zapcore.Core
	logLevel := getLogLevel(level)
	//app 层面日志
	if bapp {
		// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.WarnLevel && lvl >= logLevel
		})

		warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel && lvl >= logLevel
		})
		// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
		infoWriter := getWriter("./logs/" + filename + "/info")
		warnWriter := getWriter("./logs/" + filename + "/error")

		// 最后创建具体的Logger
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		)
	} else {
		logWriter := getWriter("./logs/" + filename)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(logWriter), logLevel),
		)
	}

	return zap.New(core)
}

//日志时间格式化
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1) + "-%Y%m%d.log", // 日志文件名格式
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		//rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)

	if err != nil {
		panic(err)
	}
	return hook
}

//LogSync flush log cache to file
func LogSync() {
	flushAppLog()
	flushInternalLog()
}

func trimmedPath(file string) string {

	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}
	// Find the penultimate separator.
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return file
	}
	return file[idx+1:]
}
