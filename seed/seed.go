package seed

import (
	"gin-mvc/configs"
	"gin-mvc/migrations"
	"gin-mvc/models"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/labstack/gommon/log"
)

func Load() {
	db := configs.GetDB()

	db.DropTableIfExists("articles", "categories", "migrations")
	migrations.Migrate(db)
	log.Info("Create Categories...")

	numOfCategory := 20

	categories := make([]models.Category, 0, numOfCategory)

	for i := 1; i < numOfCategory; i++ {
		category := models.Category{
			Name: faker.Word(),
			Desc: faker.Paragraph(),
		}

		db.Create(&category)

		categories = append(categories, category)
	}

	log.Info("Create Article...")

	numOfArticle := 20

	articles := make([]models.Article, 0, numOfArticle)

	for i := 1; i < numOfArticle; i++ {
		article := models.Article{
			Title:      faker.Sentence(),
			Excerpt:    faker.Sentence(),
			Body:       faker.Paragraph(),
			Image:      "https://source.unsplash.com/random/300x200?" + strconv.Itoa(i),
			CategoryID: uint(rand.Intn(numOfArticle) + 1),
		}

		db.Create(&article)

		articles = append(articles, article)
	}
}
