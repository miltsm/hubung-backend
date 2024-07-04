package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miltsm/hubung-service/internal/handler"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pkg/templates/*")
	r.GET("/hubung/:id", handler.New().Profile)
	r.Run(":8080")
}