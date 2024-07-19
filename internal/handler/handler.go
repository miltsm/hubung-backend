package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/miltsm/hubung-service/internal/repository"
	"github.com/miltsm/hubung-service/internal/types"
	"github.com/miltsm/hubung-service/internal/utils"
	view "github.com/miltsm/hubung-service/pkg/templates"
)

type Handler interface {
	RenderLogin(http.ResponseWriter, *http.Request)
	HandleRequestAuthChallenge(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	wauth *webauthn.WebAuthn
	repo  repository.Repository
}

func New() Handler {
	wconfig := &webauthn.Config{
		RPDisplayName: "Hubung RP service",
		RPID:          "localhost:3000",
		RPOrigins: []string{
			"http:localhost:3000",
		},
	}
	wauth, err := webauthn.New(wconfig)
	if err != nil {
		panic(err)
	}
	repo := repository.New()
	return &handler{wauth: wauth, repo: repo}
}

func (h *handler) RenderLogin(w http.ResponseWriter, r *http.Request) {
	err := view.Layout(view.Login()).Render(r.Context(), w)
	if err != nil {
		fmt.Print(err)
	}
}

func (h *handler) HandleRequestAuthChallenge(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	displayName := r.FormValue("display_name")

	// Check if name(webauthn name can be email, username or mobile no) is provided
	if len(name) == 0 {
		err := types.RequireEmailError{}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user
	user := h.repo.GetUser(name)

	// Create new user, if new user
	if user == nil {
		var err error
		if len(displayName) == 0 {
			displayName, err = utils.SubstringEmailPreAlias(name)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err = h.repo.CreateNewUser(name, displayName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var loginErr = &types.LoginError{}
	var u webauthn.User = &types.User{
		Id:          []byte(user.UserID),
		Name:        user.Email,
		DisplayName: user.Username,
	}

	opts, _, error := h.wauth.BeginRegistration(u)
	if error != nil {
		http.Error(w, loginErr.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(opts)

	// Return server challenge
	enc := json.NewEncoder(w)
	err := enc.Encode(opts)
	if err != nil {
		http.Error(w, loginErr.Error(), http.StatusInternalServerError)
		return
	}
}
