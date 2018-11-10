/**
*FileName: pkg
*Create on 2018/11/4 上午9:53
*Create by mok
*/

package conf

import (
	"os"
	"github.com/sirupsen/logrus"
)

func initLog(){
	setFile()
	setLevel()
	setFormat()
}

//设置日志文件位置
func setFile(){
	if LogFile == ""{
		LogFile = "./log/log.dat"
	}
	file,err := os.OpenFile(LogFile,os.O_APPEND | os.O_CREATE | os.O_WRONLY,0666)
	if err != nil{
		panic(err)
	}
	logrus.SetOutput(file)
}

//设置日志文件级别
func setLevel(){
	switch LogLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}

//设置日志文件格式
func setFormat(){
	if LogFormat == "text"{
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
