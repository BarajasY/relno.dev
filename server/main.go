package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Everything working OK",
		})
	})
	router.POST("", func(ctx *gin.Context) {
		r, err := ioutil.ReadAll(ctx.Request.Body)
		if err == nil {
			ctx.JSON(200, gin.H{
				"message": r,
			})
		}
	})

	router.Run()
}
