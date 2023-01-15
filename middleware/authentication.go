package middleware

import (
	"fmt"
	"instargram/config"
	"instargram/models"
	"log"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignIn struct {
	Email       string `json:"email,omitempty" binding:"omitempty,required,email"`
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"omitempty,required,len=10,number"`
	Password    string `json:"password" binding:"required"`
}

var identityKey = "sub"

func Authenticate() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		// secret key
		Key: []byte(os.Getenv("SECRET_KEY")),

		IdentityKey: identityKey,

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		IdentityHandler: func(ctx *gin.Context) interface{} {
			var user models.User
			claim := jwt.ExtractClaims(ctx)
			id := claim[identityKey]

			db := config.GetDB()
			fmt.Printf("pok-id => %#v", id)
			if db.First(&user, "id = ?", id).RowsAffected == 0 {
				return nil
			}
			fmt.Printf("user => %+v", user)
			return &user
		},

		// login => user
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var form SignIn
			var user models.User

			if err := c.ShouldBindJSON(&form); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			fmt.Printf("form => %+v", form)
			db := config.GetDB()

			if form.Email != "" {
				if db.First(&user, "email = ?", form.Email).RowsAffected == 0 {
					return nil, jwt.ErrFailedAuthentication
				}
			} else if form.PhoneNumber != "" {
				if db.First(&user, "phone_number = ?", form.PhoneNumber).RowsAffected == 0 {
					return nil, jwt.ErrFailedAuthentication
				}
			} else {
				return nil, jwt.ErrFailedAuthentication
			}

			fmt.Printf("user => %+v", user)

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		// user => payload(sub) => JWT
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				claims := jwt.MapClaims{
					identityKey: v.ID,
				}

				return claims
			}

			return jwt.MapClaims{}
		},

		// handle error
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error": message,
			})
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
