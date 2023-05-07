/**
 * @author  simple
 * @date  2023/5/7 19:08
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.LoadHTMLGlob("./templates/*")
	router.Run(":8000")

}
