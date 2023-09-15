package router

import (
	"github.com/BaiYiZi/mock-api/controller"
	"github.com/BaiYiZi/mock-api/middlewares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(middlewares.GetMethod())
	r.Use(middlewares.GetURL())
	r.Use(middlewares.GetParams())

	r.Any("/:path", controller.Control)
}
