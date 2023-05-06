/**
 * @author  simple
 * @date  2023/5/5 23:09
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8000") //默认是0.0.0.0:8080不过可以改.

}
