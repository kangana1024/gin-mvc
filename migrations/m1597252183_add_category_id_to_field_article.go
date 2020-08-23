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
			err := tx.AutoMigrate(&models.Article{}).Error

			var articles []models.Article
			tx.Find(&articles)
			for _, article := range articles {
				article.CategoryID = 2
				tx.Save(&article)
			}
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.Article{}).DropColumn("category_id").Error
		},
	}
}
