package middleware

import (
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// url := c.Request.URL
		err := utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "UnAuthorized")
			c.Abort()
			return
		}

		c.Next()
	}

}
