package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Level logrus.Level

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func SetLevel(level Level) {
	switch level {
	case PanicLevel:
		logrus.SetLevel(logrus.PanicLevel)
	case FatalLevel:
		logrus.SetLevel(logrus.FatalLevel)
	case ErrorLevel:
		logrus.SetLevel(logrus.ErrorLevel)
	case WarnLevel:
		logrus.SetLevel(logrus.WarnLevel)
	case InfoLevel:
		logrus.SetLevel(logrus.InfoLevel)
	case DebugLevel:
		logrus.SetLevel(logrus.DebugLevel)
	case TraceLevel:
		logrus.SetLevel(logrus.TraceLevel)
	}
}

func Panic(msg ...interface{}) {
	logrus.Panic(msg...)
}

func Fatal(msg ...interface{}) {
	logrus.Fatal(msg)
}

func Error(msg ...interface{}) {
	logrus.Error(msg)
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg)
}

func Info(msg ...interface{}) {
	logrus.Info(msg)
}

func Debug(msg ...interface{}) {
	logrus.Debug(msg)
}

func Trace(msg ...interface{}) {
	logrus.Trace(msg)
}

func InitLogToFile() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("logs/pluto-server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		logrus.SetOutput(fileAndStdoutWriter)
	} else {
		logrus.Info("failed to log to file.")
	}
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}
