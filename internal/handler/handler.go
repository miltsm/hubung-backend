package handler

import (
	"fmt"
	"net/http"

	"github.com/miltsm/hubung-service/internal/repository"
	view "github.com/miltsm/hubung-service/pkg/templates"
)

type Handler interface {
	RenderLogin(http.ResponseWriter, *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo repository.Repository
}

func New() Handler {
	repo := repository.New()
	return &handler{ repo: repo }
}

func (h *handler) RenderLogin(w http.ResponseWriter, r *http.Request) {
	err := view.Layout(view.Login()).Render(r.Context(), w)
	if err != nil {
		fmt.Print(err)
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	//TODO: passwordless login
}