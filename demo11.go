/**
 * @author  simple
 * @date  2023/5/8 10:15
 * @version 1.0
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, err := c.Cookie("Username")

		if err != nil {
			cookie = "NotSet"
			//名称, 值,有效时长,有效路径,域,安全模式是否开启,httponly是否开启
			c.SetCookie("Username", "for no one", 60*60, "/", "127.0.0.1", false, true)

		}

		fmt.Printf("Cookie value: %s \n", cookie)
		c.String(200, "test cookie")
	})

	router.Run(":8000")
}
