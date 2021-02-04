package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
	"taylors/utils"
)

var errorLogger *log.Logger

func Init(path string, is_dev bool, level string) {
	if is_dev {
		init_dev()
	} else {
		init_standard(path, level)
	}
}

func init_dev() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	errorLogger = log.StandardLogger()
}

func init_standard(path, level string) {
	if !utils.FileExist(path) {
		_, err := os.Create(path)
		if err != nil {
			panic("初始化日志路径失败-创建路径失败")
		}
	}

	outPutFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	lev, err := log.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(outPutFile)
	log.SetLevel(lev)
	errorLogger = log.StandardLogger()
}

func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
