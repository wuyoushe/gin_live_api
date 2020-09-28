package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"

	"github.com/rifflock/lfshook"
	config "github.com/wuyoushe/gin_live_api/conf"
)

//日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	filePath := getLoggerFileFullPath()
	F = openLoggerFile(filePath)

	logFilePath := config.Log_FILE_PATH
	logFileName := config.LOG_FILE_NAME

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	fmt.Println(fileName)

	// 写入文件
	_, err := os.Stat(fileName)

	switch {
	case os.IsNotExist(err):
		mkNewDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日期格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05"})

	//设置rotatelogs
	logWriter, err := rotatelogs.New(
		//分割后的文件名称
		fileName+".%Y%m%d.log",

		//生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//新增Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()

		//处理请求
		c.Next()

		//结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := c.Request.Method

		//请求路由
		reqUri := c.Request.RequestURI

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

//日志记录到MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

///日志记录到ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//日志记录到MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func mkNewDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLoggerFilePath(), os.ModePerm)

	if err != nil {
		panic(err)
	}
}

func getLoggerFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}
func getLoggerFileFullPath() string {
	prefixPath := getLoggerFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLoggerFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkNewDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}
