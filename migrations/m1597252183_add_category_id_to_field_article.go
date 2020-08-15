package migrations

import (
	"gin-mvc/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1597252183AddCategoryIDToFieldArticle() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1597252183",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("articles").Error
		},
	}
}
