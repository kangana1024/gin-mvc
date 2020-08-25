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
	authenticate := middleware.Authenticate().MiddlewareFunc()

	authorize := middleware.Authorize()

	userController := controllers.Users{DB: db}
	userGroup := v1.Group("users")
	userGroup.Use(authenticate, authorize)
	{
		userGroup.GET("/", userController.FindAll)

		userGroup.GET(":id", userController.FindOne)
		userGroup.PATCH(":id", userController.UpdateByID)
		userGroup.DELETE(":id", userController.Delete)
		userGroup.PATCH(":id/promote", userController.Promote)
		userGroup.PATCH(":id/demote", userController.Demote)

		userGroup.POST("/", userController.Create)
	}

	authGroup := v1.Group("/auth")

	authController := controllers.Auth{DB: db}

	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
		authGroup.GET("/profile", authenticate, authController.GetProfile)
		authGroup.PATCH("/profile", authenticate, authController.UpdateProfile)
	}

	articlesController := controllers.Article{
		DB: db,
	}
	articlesGroup := v1.Group("/articles")
	articlesGroup.GET("/", articlesController.FindAll)

	articlesGroup.GET(":id", articlesController.FindOne)
	articlesGroup.Use(authenticate, authorize)
	{
		articlesGroup.PATCH(":id", articlesController.UpdateByID)
		articlesGroup.DELETE(":id", articlesController.DeleteByID)

		articlesGroup.POST("/", articlesController.Create)

	}
	categoryController := controllers.Category{
		DB: db,
	}
	categoryGroup := v1.Group("categories")
	categoryGroup.Use(authenticate, authorize)
	{
		categoryGroup.GET("/", categoryController.FindAll)

		categoryGroup.GET(":id", categoryController.FindOne)
		categoryGroup.PATCH(":id", categoryController.UpdateByID)
		categoryGroup.DELETE(":id", categoryController.DeleteByID)

		categoryGroup.POST("/", categoryController.Create)

	}

	return
}
