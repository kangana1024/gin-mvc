package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type article struct {
	ID uint
	Title string
	Body string
}

func Serve(r *gin.Engine){
	articles := []article{
		article{
			ID: 1,
			Title: "Title#1",
			Body: "Body#1",
		},
		article{
			ID: 2,
			Title: "Title#2",
			Body: "Body#2",
		},
		article{
			ID: 3,
			Title: "Title#3",
			Body: "Body#3",
		},
		article{
			ID: 4,
			Title: "Title#4",
			Body: "Body#4",
		},
		article{
			ID: 5,
			Title: "Title#5",
			Body: "Body#5",
		},
	}
	articlesGroup :=	r.Group("api/v1/articles")

	articlesGroup.GET("/", func(ctx *gin.Context){
		result := articles
		if limit := ctx.Query("limit"); limit != "" {
			n,_ := strconv.Atoi(limit)

			result = result[:n]
		}
		ctx.JSON(http.StatusOK, gin.H{"ariticles":result})
	})

	articlesGroup.GET(":id", func(ctx *gin.Context){
		id, _ := strconv.Atoi(ctx.Param("id"))
		for _,item := range articles {
			if item.ID == uint(id){
				ctx.JSON(http.StatusOK,gin.H{"article": item})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{"error":"Article not found"})
	})

	
}