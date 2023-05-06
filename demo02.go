/**
 * @author  simple
 * @date  2023/5/5 23:29
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func someGet(c *gin.Context) {
	println("someget")
	c.Writer.WriteString("get 方法")
}

func somePost(c *gin.Context) {
	println("somepost")
	c.Writer.WriteString("POST 方法")
}

func main() {
	r := gin.Default()
	r.GET("/get", someGet)
	r.POST("/post", somePost)
	r.Run(":8000")
}
