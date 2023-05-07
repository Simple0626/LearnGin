/**
 * @author  simple
 * @date  2023/5/7 17:17
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	/*	<form action="/search?page=1">
		<input name="key" >
		<input type="submit">
		</form>*/
	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.POST("/search", search)
	router.Run(":8000")
}

func search(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "0")
	key := ctx.PostForm("key")
	hobby := ctx.PostFormArray("hobby")
	ctx.String(200, "this is search page: %s key: %s hobby: %s ", page, key, hobby)
}
