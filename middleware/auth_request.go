package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cledupe/jwt-auth/infrastructure"
	"github.com/cledupe/jwt-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func TokenAuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")

	if err != nil {
		log.Println("cookie not exist")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		log.Println("parsing token error")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			log.Println("token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		infrastructure.DB.Db.First(&user, claims["user"])

		if user.ID == 0 {
			log.Println("user not found")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()

	} else {
		log.Println("token not valid")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
