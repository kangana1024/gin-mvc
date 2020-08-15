package routes

import (
	"gin-mvc/configs"
	"gin-mvc/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := configs.GetDB()
	articlesController := controllers.Article{
		DB: db,
	}
	articlesGroup := r.Group("api/v1/articles")
	{
		articlesGroup.GET("/", articlesController.FindAll)

		articlesGroup.GET(":id", articlesController.FindOne)
		articlesGroup.PATCH(":id", articlesController.UpdateByID)
		articlesGroup.DELETE(":id", articlesController.DeleteByID)

		articlesGroup.POST("/", articlesController.Create)

	}
	categoryController := controllers.Category{
		DB: db,
	}
	categoryGroup := r.Group("api/v1/categories")
	{
		categoryGroup.GET("/", categoryController.FindAll)

		categoryGroup.GET(":id", categoryController.FindOne)
		categoryGroup.PATCH(":id", categoryController.UpdateByID)
		categoryGroup.DELETE(":id", categoryController.DeleteByID)

		categoryGroup.POST("/", categoryController.Create)

	}

	return
}
