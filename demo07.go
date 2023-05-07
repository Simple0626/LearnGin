/**
 * @author  simple
 * @date  2023/5/7 18:44
 * @version 1.0
 */

package main

import (
	"LearnGin/model"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.GET("/register", toregister)
	router.POST("/register", register)
	router.Run(":8000")
}

func toregister(ctx *gin.Context) {
	ctx.HTML(200, "register.html", nil)

}

func register(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBind(&user)
	ctx.String(200, "获取到 %s", user)
}
