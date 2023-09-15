package middlewares

import (
	"github.com/gin-gonic/gin"
)

func GetURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("req_method", c.Request.Method)
	}
}
