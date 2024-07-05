package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/miltsm/hubung-service/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.LoadHTMLGlob("pkg/templates/*")
	h := handler.New()
	r.GET("/hubung/:id", h.Profile)
	r.GET("/login", h.Profile)
	r.Run(":8080")
}