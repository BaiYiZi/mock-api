package middlewares

import (
	"github.com/gin-gonic/gin"
)

func GetMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("req_path", c.Param("path"))
	}
}
