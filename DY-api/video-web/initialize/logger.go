/*
 * @Date: 2023-01-19 11:21:47
 * @LastEditors: zhang zhao
 * @LastEditTime: 2023-01-22 18:26:08
 * @FilePath: /simple-DY/DY-api/video-web/initialize/logger.go
 * @Description: zap日志框架配置
 */
package initialize

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {

	// 设置一些基本日志格式
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// // 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	// infoWriter := getWriter("./logs/info.log")
	// errorWriter := getWriter("./logs/error.log")

	logWriter := getWriter("./logs/video-web.log")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel), //打印到控制台
		zapcore.NewCore(encoder, zapcore.AddSync(logWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(logWriter), errorLevel),
	)

	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	// errorLogger = logger.Sugar()
	zap.ReplaceGlobals(logger)
}

// 生成 rotatelogs 的 Logger
// 保存7天内的日志，每1小时(整点)分割一次日志
func getWriter(filename string) io.Writer {

	hook, err := rotatelogs.New(
		filename[:len(filename)-4]+"_%Y_%m_%d_%H.log", // 日期时间格式化
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// }
// func Debug(args ...interface{}) {
// 	errorLogger.Debug(args...)
// }

// func Debugf(template string, args ...interface{}) {
// 	errorLogger.Debugf(template, args...)
// }

// func Info(args ...interface{}) {
// 	errorLogger.Info(args...)
// }

// func Infof(template string, args ...interface{}) {
// 	errorLogger.Infof(template, args...)
// }

// func Warn(args ...interface{}) {
// 	errorLogger.Warn(args...)
// }

// func Warnf(template string, args ...interface{}) {
// 	errorLogger.Warnf(template, args...)
// }

// func Error(args ...interface{}) {
// 	errorLogger.Error(args...)
// }

// func Errorf(template string, args ...interface{}) {
// 	errorLogger.Errorf(template, args...)
// }

// func DPanic(args ...interface{}) {
// 	errorLogger.DPanic(args...)
// }

// func DPanicf(template string, args ...interface{}) {
// 	errorLogger.DPanicf(template, args...)
// }

// func Panic(args ...interface{}) {
// 	errorLogger.Panic(args...)
// }

// func Panicf(template string, args ...interface{}) {
// 	errorLogger.Panicf(template, args...)
// }

// func Fatal(args ...interface{}) {
// 	errorLogger.Fatal(args...)
// }

// func Fatalf(template string, args ...interface{}) {
// 	errorLogger.Fatalf(template, args...)
// }
