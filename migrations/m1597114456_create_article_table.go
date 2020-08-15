package migrations

import (
	"gin-mvc/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1597114456CreateArticleTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1597114456",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(&models.Article{}).Error

			var articles []models.Article
			tx.Find(&articles)
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("articles").Error
		},
	}
}
