package migrations

import (
	"gin-mvc/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1597250312CreateCategoryTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1597250312",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Category{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("categories").Error
		},
	}
}
