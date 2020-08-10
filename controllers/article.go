package controllers

import (
	"gin-mvc/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateArticalForm struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
type Article struct {
}

var articles []models.Article = []models.Article{
	{
		ID:    1,
		Title: "Title#1",
		Body:  "Body#1",
	},
	{
		ID:    2,
		Title: "Title#2",
		Body:  "Body#2",
	},
	{
		ID:    3,
		Title: "Title#3",
		Body:  "Body#3",
	},
	{
		ID:    4,
		Title: "Title#4",
		Body:  "Body#4",
	},
	{
		ID:    5,
		Title: "Title#5",
		Body:  "Body#5",
	},
}

func (a *Article) FindAll(ctx *gin.Context) {
	result := articles
	if limit := ctx.Query("limit"); limit != "" {
		n, _ := strconv.Atoi(limit)

		result = result[:n]
	}
	ctx.JSON(http.StatusOK, gin.H{"ariticles": result})
}
func (a *Article) FindOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for _, item := range articles {
		if item.ID == uint(id) {
			ctx.JSON(http.StatusOK, gin.H{"article": item})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
}

func (a *Article) Create(ctx *gin.Context) {
	var form CreateArticalForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		ID:    uint(len(articles)) + 1,
		Title: form.Title,
		Body:  form.Body,
		Image: "",
	}

	// Get File
	file, _ := ctx.FormFile("image")

	// Create Path
	// ID => 8, upload/articles/8/image.png
	path := "uploads/articles/" + strconv.Itoa(int(article.ID))

	os.MkdirAll(path, 0755)

	// Upload file
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "file": filename})
	}

	// Attach File to article
	article.Image = os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + filename

	articles = append(articles, article)

	ctx.JSON(http.StatusCreated, gin.H{"article": article})
}
