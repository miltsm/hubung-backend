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
	contentType := c.ContentType()
	if (contentType == "application/json") {
		c.IndentedJSON(http.StatusOK, profile) 
	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{ 
			"name": profile.Name,
			"link_name": profile.Categories[0].Hubungan[0].Name,
			"link": profile.Categories[0].Hubungan[0].Link,
		})
	}
} 