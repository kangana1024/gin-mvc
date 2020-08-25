package controllers

import (
	"gin-mvc/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type createArticalForm struct {
	Title   string                `form:"title" binding:"required"`
	Body    string                `form:"body" binding:"required"`
	Excerpt string                `form:"excerpt" binding:"required"`
	Image   *multipart.FileHeader `form:"image" binding:"required"`
}
type updateArticalForm struct {
	Title   string                `form:"title"`
	Body    string                `form:"body"`
	Excerpt string                `form:"excerpt"`
	Image   *multipart.FileHeader `form:"image"`
}
type articleResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Excerpt    string `json:"excerpt"`
	Body       string `json:"body"`
	Image      string `json:"image"`
	CategoryID uint   `json:"categoryId"`
	Category   struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	User struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"user"`
}
type Article struct {
	DB *gorm.DB
}
type articlePaging struct {
	Items  []articleResponse `json:"items"`
	Paging *PagingResult     `json:"paging"`
}

func (a *Article) FindAll(ctx *gin.Context) {
	var articles []models.Article

	pagination := pagination{
		ctx:     ctx,
		query:   a.DB.Preload("User").Preload("Category").Order("id desc"),
		records: &articles,
	}
	paging := pagination.paginate()
	var serializedArticles []articleResponse
	copier.Copy(&serializedArticles, &articles)
	ctx.JSON(http.StatusOK, gin.H{"articles": articlePaging{
		Items:  serializedArticles,
		Paging: paging,
	}})
}
func (a *Article) DeleteByID(ctx *gin.Context) {
	article, err := a.findArticleByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	a.DB.Delete(&article)
	// a.DB.Unscoped().Delete(&article)
	ctx.Status(http.StatusNoContent)
}
func (a *Article) UpdateByID(ctx *gin.Context) {
	var form updateArticalForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	article, err := a.findArticleByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := a.DB.Model(&article).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	a.setArticleImage(ctx, article)

	var serializedArticle articleResponse

	copier.Copy(&serializedArticle, article)
	ctx.JSON(http.StatusOK, gin.H{"article": serializedArticle})

}
func (a *Article) FindOne(ctx *gin.Context) {
	article, err := a.findArticleByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	serializedArticle := articleResponse{}
	copier.Copy(&serializedArticle, &article)
	ctx.JSON(http.StatusOK, gin.H{"article": serializedArticle})
}

func (a *Article) Create(ctx *gin.Context) {
	var form createArticalForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	var article models.Article
	user, _ := ctx.Get("sub")
	article.User = *user.(*models.User)

	copier.Copy(&article, &form)

	if err := a.DB.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if err := a.setArticleImage(ctx, &article); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedArticle := articleResponse{}
	copier.Copy(&serializedArticle, &article)
	ctx.JSON(http.StatusCreated, gin.H{"article": serializedArticle})
}

func (a *Article) setArticleImage(ctx *gin.Context, article *models.Article) error {
	file, err := ctx.FormFile("image")
	if err != nil || file == nil {
		return err
	}

	if article.Image != "" {

		article.Image = strings.Replace(article.Image, os.Getenv("HOST"), "", 1)

		pwd, _ := os.Getwd()

		os.Remove(pwd + article.Image)

	}

	path := "uploads/articles/" + strconv.Itoa(int(article.ID))
	os.MkdirAll(path, 0755)
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		return err
	}

	article.Image = os.Getenv("HOST") + "/" + filename

	a.DB.Save(article)

	return nil
}
func (a *Article) findArticleByID(ctx *gin.Context) (*models.Article, error) {
	var article models.Article

	id := ctx.Param("id")

	if err := a.DB.Preload("User").Preload("Category").First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}
