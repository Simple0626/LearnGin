/**
 * @author  simple
 * @date  2023/5/7 19:47
 * @version 1.0
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	ginServer := gin.Default()
	// gin.Accounts 是 map[string]string 的一种快捷方式
	routergroup := ginServer.Group("/admin", gin.BasicAuth(gin.Accounts{
		//从数据库查
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	routergroup.GET("/secrets", Handlera)
	ginServer.Run(":8000")
}

func Handlera(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}
