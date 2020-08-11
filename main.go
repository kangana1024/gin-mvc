package main

import (
	"gin-mvc/configs"
	"gin-mvc/migrations"
	"gin-mvc/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Env")
	}

	configs.InitDB()
	defer configs.CloseDB()
	db := configs.GetDB()
	migrations.Migrate(db)

	r := gin.Default()
	r.Static("/uploads", "./uploads")
	os.MkdirAll("uploads", 0755)
	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}
	url := ginSwagger.URL(os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/swagger/doc.json") // The url pointing to API definition
	routes.Serve(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err = r.Run(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Run Fail")
	}

}
