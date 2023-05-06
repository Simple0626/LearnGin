/**
 * @author  simple
 * @date  2023/5/6 23:42
 * @version 1.0
 */

package main

import (
	"github.com/gin-gonic/gin"
)

func gethtml(c *gin.Context) {
	username := c.Query("username")
	println(username)
	c.HTML(200, "index.html", gin.H{
		"username": username,
	})
}

func main() {
	router := gin.Default()
	//导入模板:
	router.LoadHTMLGlob("./templates/*")
	router.GET("/demo05", gethtml)
	router.Run(":8000")
}
