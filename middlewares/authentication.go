package middlewares

import (
	"jwt-go/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userData, err := helpers.VerifyToken(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})

			return
		} else {
			c.Set("userData", userData)
			c.Next()
		}
	}
}
