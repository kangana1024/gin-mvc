package middleware

import (
	"gin-mvc/models"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.Get("sub")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		enforcer := casbin.NewEnforcer("configs/acl_model.conf", "configs/policy.csv")

		ok = enforcer.Enforce(user.(*models.User), ctx.Request.URL.Path, ctx.Request.Method)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}
