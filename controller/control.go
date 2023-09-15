package controller

import (
	"github.com/BaiYiZi/mock-api/services"
	"github.com/gin-gonic/gin"
)

func Control(c *gin.Context) {
	res, code, err := services.Query(c)
	if err != nil {
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(code, res)
}
