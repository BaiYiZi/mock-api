package main

import (
	"github.com/BaiYiZi/mock-api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)

	r.Run(":9090")
}
