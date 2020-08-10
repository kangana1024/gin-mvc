package routes

import (
	"gin-mvc/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {

	articlesController := controllers.Article{}
	articlesGroup := r.Group("api/v1/articles")
	{
		articlesGroup.GET("/", articlesController.FindAll)

		articlesGroup.GET(":id", articlesController.FindOne)

		articlesGroup.POST("/", articlesController.Create)
	}

	return
}
