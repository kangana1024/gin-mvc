package controllers

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CreateArticalForm struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
type Article struct {
}

func (a *Article) FindAll(ctx *gin.Context) {
}
func (a *Article) FindOne(ctx *gin.Context) {
}

func (a *Article) Create(ctx *gin.Context) {
}
