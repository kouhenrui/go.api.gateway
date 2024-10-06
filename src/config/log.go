package config

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

// LogConfig holds the configuration for the logger
type LogConfig struct {
	LogPath  string // 日志文件存放路径
	LogLevel string // 日志级别
}

// InitLogger initializes the logger with rotation and custom log formats
func InitLogger(logConf *LogConfig) {
	//确定文件是否存在
	if _, err := os.Stat(logConf.LogPath); os.IsNotExist(err) {
		if err := os.MkdirAll(logConf.LogPath, 0755); err != nil {
			panic(fmt.Sprintf("日志文件存放地址创建错误: %s", err))
		}
	}

	// Set log file path with rotation by day
	logWriter, err := rotatelogs.New(
		logConf.LogPath+"/%Y%m%d.log",             // 日志文件名格式
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 日志保留7天
		rotatelogs.WithRotationTime(24*time.Hour), // 每24小时切割一次日志
	)
	if err != nil {
		panic(fmt.Sprintf("日志文件创建错误: %s", err))
	}

	// Set log level
	level, err := logrus.ParseLevel(logConf.LogLevel)
	if err != nil {
		panic(fmt.Sprintf("日志级别解析错误: %s", err))
	}
	Logger.SetLevel(level)

	// Define custom formatters for info and error logs
	logFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 设置时间格式
		PrettyPrint:     false,                 // 关闭 Pretty Print 保证日志是紧凑的 JSON 格式
	}

	// Map log levels to the log writer
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// Create hooks for info and error logs
	hook := lfshook.NewHook(writeMap, logrus.Formatter(logFormatter))

	// Add hooks to the logger
	Logger.AddHook(hook)

	fmt.Println("日志初始化成功")
}
