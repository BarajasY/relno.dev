package main

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	posts "relnodev/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

//Forms DbPool and adds it to gin's context.
func MakePool() *sql.DB {
	conn_string := os.Getenv("CONN_STRING")
	db, err := sql.Open("pgx", conn_string)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create connection pool \n %s", err)
		os.Exit(1)
	}

	return db
}

func main() {
	godotenv.Load()

	router := gin.Default()
	db := MakePool()

	//CORS POLICY
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
	router.GET("/posts", func(ctx *gin.Context) {
		posts.ShowAvailablePosts(ctx, db)
	})
	router.GET("/tags", func(ctx *gin.Context) {
		posts.RetrievePopularTags(ctx, db)
	})
	router.GET("/posts/:title", func(ctx *gin.Context) {
		title := ctx.Param("title")
		posts.ShowOnePost(ctx, db, title)
	})

	router.Run()
}
