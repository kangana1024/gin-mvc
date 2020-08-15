package controllers

import (
	"gin-mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type createCategoryForm struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}
type categoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
type updateCategoryResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
type Category struct {
	DB *gorm.DB
}
type categoryPaging struct {
	Items  []categoryResponse `json:"items"`
	Paging *PagingResult      `json:"paging"`
}

func (c *Category) FindAll(ctx *gin.Context) {
	var categories []models.Category

	pagination := pagination{
		ctx:     ctx,
		query:   c.DB.Order("id desc"),
		records: &categories,
	}
	paging := pagination.paginate()
	var serializedCategories []categoryResponse
	copier.Copy(&serializedCategories, &categories)
	ctx.JSON(http.StatusOK, gin.H{"categories": categoryPaging{
		Items:  serializedCategories,
		Paging: paging,
	}})
}
func (c *Category) DeleteByID(ctx *gin.Context) {
	category, err := c.findCategoryByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.DB.Delete(&category)
	// c.DB.Unscoped().Delete(&article)
	ctx.Status(http.StatusNoContent)
}
func (c *Category) UpdateByID(ctx *gin.Context) {
	var form updateCategoryResponse

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	category, err := c.findCategoryByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&category).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory categoryResponse

	copier.Copy(&serializedCategory, category)
	ctx.JSON(http.StatusOK, gin.H{"category": serializedCategory})

}
func (c *Category) FindOne(ctx *gin.Context) {
	category, err := c.findCategoryByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	serializedCategory := categoryResponse{}
	copier.Copy(&serializedCategory, &category)
	ctx.JSON(http.StatusOK, gin.H{"category": serializedCategory})
}

func (c *Category) Create(ctx *gin.Context) {
	var form createCategoryForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	var category models.Category

	copier.Copy(&category, &form)

	if err := c.DB.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedCategory := categoryResponse{}
	copier.Copy(&serializedCategory, &category)
	ctx.JSON(http.StatusCreated, gin.H{"category": serializedCategory})
}
func (c *Category) findCategoryByID(ctx *gin.Context) (*models.Category, error) {
	var category models.Category

	id := ctx.Param("id")

	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}
