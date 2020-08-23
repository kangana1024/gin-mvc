package migrations

import (
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) {
	// db := configs.GetDB()
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1597114456CreateArticleTable(),
			m1597250312CreateCategoryTable(),
			m1597252183AddCategoryIDToFieldArticle(),
			m1597483362AddCategoryIDToFieldArticle(),
			m1598019416CreateUserTable(),
		},
	)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
