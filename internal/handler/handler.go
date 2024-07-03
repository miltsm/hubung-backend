package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miltsm/hubung-service/internal/repository"
)

type Handler interface {
	Profile(c *gin.Context)
}

type handler struct {
	repo repository.Repository
}

func New() Handler {
	repo := repository.New()
	return &handler{ repo: repo }
}

func (h *handler) Profile(c *gin.Context) {
	id := c.Param("id")
	profile := h.repo.GetProfile(id)
	c.IndentedJSON(http.StatusOK, profile)  
} 