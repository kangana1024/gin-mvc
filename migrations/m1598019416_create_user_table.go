package migrations

import (
	"gin-mvc/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1598019416CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1598019416",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	}
}
