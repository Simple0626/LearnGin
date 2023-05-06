/**
 * @author  simple
 * @date  2023/5/6 23:30
 * @version 1.0
 */

package main

import "github.com/gin-gonic/gin"

func getParse(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.String(200, "name is %s age is %s", name, age)
}

func main() {
	router := gin.Default()
	router.GET("/demo04/:name/:age", getParse)

	router.Run(":8000")

}
