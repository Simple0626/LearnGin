/**
 * @author  simple
 * @date  2023/5/6 15:48
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func getPostval(c *gin.Context) {
	username := c.PostForm("name")
	passwd := c.DefaultPostForm("passwd", "xxxxxxxx")
	c.String(200, "获取到的name = %s,paswd = %s ", username, passwd)
}

func main() {
	router := gin.Default()
	router.GET("/demo03get", func(c *gin.Context) {
		username := c.Query("name")
		passwd := c.DefaultQuery("passwd", "xxxxxx") //第二个是默认值.
		c.String(200, "this is getkey() name= %s  passwd = %s ", username, passwd)
	})
	router.POST("/demo03post", getPostval)
	router.Run(":8000")
}
