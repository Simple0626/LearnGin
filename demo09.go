/**
 * @author  simple
 * @date  2023/5/7 19:33
 * @version 1.0
 */

package main

import (
	"github.com/gin-gonic/gin"
)

/*
中间件
*/

func MyHandler(ctx *gin.Context) {
	println("这里是中间件")
}
func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//完整路径名
		path := c.FullPath()
		method := c.Request.Method
		println("path : ", path, "method", method)

	}
}

func main() {
	router := gin.Default()
	//router.Use(MyHandler, Handler())
	router.GET("/index", MyHandler, func(c *gin.Context) {
		c.String(200, "index page")
	})
	router.Run(":8000")
}
