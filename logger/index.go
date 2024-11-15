package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zL *zap.Logger

func init() {
	// 创建文件写入器
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 初始化 Zap 日志
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zL = zap.New(zapcore.NewCore(
		// zapcore.NewJSONEncoder(encoderConfig), // JOSN 格式输出日志
		zapcore.NewConsoleEncoder(encoderConfig), // 常规的控制台输出格式
		// zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.AddSync(file), // 输出到文件
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
	))
}

func GetLogger() *zap.Logger {
	return zL
}

func Sync() {
	GetLogger().Sync()
}

func Infof(template string, args ...interface{}) {
	zL.Sugar().Infof(template, args...)
}

func Debugf(template string, args ...interface{}) {
	zL.Sugar().Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	zL.Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zL.Sugar().Errorf(template, args...)
}
