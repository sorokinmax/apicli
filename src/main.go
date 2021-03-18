package main

import (
	"flag"
	"runtime"
	"strconv"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	initPort = flag.Int("port", 8001, "Listening port for service")
	log      = logrus.New()
	version  = "dev"
	build    = ""
	commit   = ""
)

func init() {
	flag.Parse()

	log.Out = &lumberjack.Logger{
		Filename:   currentDir() + "/apns.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     30, // days
	}

	log.Formatter = &logrus.JSONFormatter{}
	//log.ReportCaller = true
	//log.SetLevel(log.InfoLevel)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultErrorWriter = &panicWriter{}
	gin.DefaultWriter = &accessWriter{}
}

func main() {
	runtime.GOMAXPROCS(1)

	router := gin.Default()

	router.GET("/ping", pingHandler)
	router.POST("/run", runHandler)

	log.Infof("Version: %s(%s:%s)", version, build, commit)
	log.Infof("API is available at :%s", strconv.Itoa(*initPort))

	router.Run(":" + strconv.Itoa(*initPort))

}
