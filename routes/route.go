package routes

import (
	"gin-mvc/configs"
	"gin-mvc/controllers"
	"gin-mvc/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := configs.GetDB()
	v1 := r.Group("/api/v1")

	authGroup := v1.Group("auth")

	authController := controllers.Auth{DB: db}

	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
	}

	articlesController := controllers.Article{
		DB: db,
	}
	articlesGroup := v1.Group("api/v1/articles")
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
	categoryGroup := v1.Group("api/v1/categories")
	{
		categoryGroup.GET("/", categoryController.FindAll)

		categoryGroup.GET(":id", categoryController.FindOne)
		categoryGroup.PATCH(":id", categoryController.UpdateByID)
		categoryGroup.DELETE(":id", categoryController.DeleteByID)

		categoryGroup.POST("/", categoryController.Create)

	}

	return
}
