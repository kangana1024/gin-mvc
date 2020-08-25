package middleware

import (
	"fmt"
	"gin-mvc/configs"
	"gin-mvc/models"
	"log"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

var identityKey = "sub"

func Authenticate() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:           []byte(os.Getenv("SECRET_KEY")),
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		IdentityHandler: func(c *gin.Context) interface{} {
			var user models.User
			claims := jwt.ExtractClaims(c)
			id := claims[identityKey]

			db := configs.GetDB()
			if db.First(&user, uint(id.(float64))).RecordNotFound() {
				return nil
			}

			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var form login

			var user models.User

			if err := c.ShouldBindJSON(&form); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			db := configs.GetDB()

			if db.Where("email=?", form.Email).First(&user).RecordNotFound() {
				return nil, jwt.ErrFailedAuthentication
			}
			fmt.Println("P1 :" + user.Password + " P2 : " + form.Password)
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				claims := jwt.MapClaims{
					identityKey: v.ID,
				}
				return claims
			}

			return jwt.MapClaims{}

		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": message,
			})
		},
	})

	if err != nil {
		log.Fatal("JWT ERROR : " + err.Error())
	}

	return authMiddleware
}
