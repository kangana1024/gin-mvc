package configs

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Env")
	}
	db, err = gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_CONN"))
	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(gin.Mode() == gin.DebugMode)

}
func GetDB() *gorm.DB {
	return db
}
func CloseDB() {
	db.Close()
}
