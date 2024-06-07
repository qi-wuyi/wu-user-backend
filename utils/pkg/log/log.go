package log

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

func LogConfig() {
	//定制输出目录
	logFilePath := "./Logs"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		hlog.Fatalf(err.Error())
		return
	}

	//将文件名设置为日期
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			hlog.Fatalf(err.Error())
			return
		}
	}
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	logger := hertzzap.NewLogger(hertzzap.WithCoreEnc(zapcore.NewConsoleEncoder(cfg)))

	//提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   //一个文件最大可达20m
		MaxBackups: 5,    // 最多可以同时保存5个文件
		MaxAge:     10,   // 一个文件最多可以保持10天
		Compress:   true, //用gzip压缩
	}
	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
}
