/**
 * @author  simple
 * @date  2023/5/9 19:37
 * @version 1.0
 */

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//日志中间件gin 自带的

func main() {
	router := gin.Default()
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	//gin.DefaultWriter  = io.MultiWriter(file,os.Stdout)  //写文件的同时还会在控制台打印
	println()
	router.Run()

}
