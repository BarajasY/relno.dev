package posts

import (
	"database/sql"
	"fmt"
	"os"
	markdown "relnodev/utils"

	"github.com/gin-gonic/gin"
)

type tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type post_tag struct {
	Post_id int `json:"Post_id"`
	Tag_id  int `json:"Tag_id"`
}

type post struct {
	//UPPERCASE IS IMPORTANT, otherwise they count as not exported
	Post_id      int    `json:"post_id"`
	Post_title   string `json:"post_title"`
	Post_date    int    `json:"post_date"`
	Post_content string `json:"post_content"`
}

type postWithTags struct {
	Post_id      int    `json:"post_id"`
	Post_title   string `json:"post_title"`
	Post_date    int    `json:"post_date"`
	Post_content string `json:"post_content"`
	Post_tags    []tag  `json:"post_tags"`
}

func AddPostToDatabase() {
	fmt.Printf("Hola")
}

func ShowAvailablePosts(ctx *gin.Context, db *sql.DB) {
	posts := []postWithTags{}

	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to select posts \n %s", err)
		os.Exit(1)
	}
	for rows.Next() {
		var data post = post{}
		tags := []tag{}
		rows.Scan(&data.Post_id, &data.Post_title, &data.Post_content, &data.Post_date)
		println(data.Post_id)
		rows, err := db.Query("Select * from post_tags where post_id = $1", data.Post_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to select post_tags \n %s", err)
			os.Exit(1)
		}
		for rows.Next() {
			var post_tag post_tag = post_tag{}
			rows.Scan(&post_tag.Post_id, &post_tag.Tag_id)
			rows, err := db.Query("Select * from tags where id = $1", post_tag.Tag_id)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unable to select tags \n %s", err)
				os.Exit(1)
			}
			for rows.Next() {
				var tag tag = tag{}
				rows.Scan(&tag.Id, &tag.Name)
				tags = append(tags, tag)
			}
		}
		var markdowndata postWithTags = postWithTags{
			Post_id:      data.Post_id,
			Post_title:   data.Post_title,
			Post_date:    data.Post_date,
			Post_content: string(markdown.MdToHTML(data.Post_content)),
			Post_tags: tags,
		}
		posts = append(posts, markdowndata)
	}

	ctx.JSON(200, posts)

}
