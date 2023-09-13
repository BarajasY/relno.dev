package main

import (
	"io"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
	}))
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Everything working OK",
		})
	})
	router.POST("", func(ctx *gin.Context) {
		r, err := io.ReadAll(ctx.Request.Body)
		if err == nil {
			ctx.JSON(200, gin.H{
				"message": r,
			})
		}
	})

	router.Run()
}
