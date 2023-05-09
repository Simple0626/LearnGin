/**
 * @author  simple
 * @date  2023/5/9 19:50
 * @version 1.0
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

/*
gin 集成logrus
*/

var (
	logFilePath = "./"
	logFileName = "system"
)

func logerMinddleware() gin.HandlerFunc {
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	//写入文件
	src, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	//实例化:
	logger := logrus.New()
	//设置日志级别:
	logger.SetLevel(logrus.DebugLevel)
	//设置输出:
	logrus.SetOutput(src)

	//设置

}
