package handler

import (
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/miltsm/hubung-service/internal/repository"
	view "github.com/miltsm/hubung-service/pkg/templates"
)

type Handler interface {
	RenderLogin(http.ResponseWriter, *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	wauth *webauthn.WebAuthn
	repo repository.Repository
}

func New() Handler {
	wconfig := &webauthn.Config{
		RPDisplayName: "Hubung RP service",
		RPID: "localhost:3000",
		RPOrigins: []string{
			"http:localhost:3000",
		},
	}
	wauth, err := webauthn.New(wconfig)
	if err != nil {
		panic(err)
	}
	repo := repository.New()
	return &handler{ wauth: wauth, repo: repo }
}

func (h *handler) RenderLogin(w http.ResponseWriter, r *http.Request) {
	err := view.Layout(view.Login()).Render(r.Context(), w)
	if err != nil {
		fmt.Print(err)
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	displayName := r.FormValue("display_name")

	if len(name) == 0 {
		http.Error(w, "Email address required", http.StatusBadRequest)
		return
	}

	if len(displayName) == 0 {
		displayName, err := substringEmailPreAlias(name); if err != nil {
			http.Error(w, "Email address required", http.StatusBadRequest)
			return
		}
		fmt.Println(displayName)
	}

	//TODO: User DB creation
	// uId, passkeyUId, publicKey := exec.Command("uuidgen"), 
	// h.wauth.BeginRegistration(user)
}

type RequireEmailError struct {}

type EmptyStringError struct {}
func (e *EmptyStringError) Error() string {
	return "Empty string error!"
}

func substringEmailPreAlias(e string) (*string, error) {
	if len(e) == 0 {
		return nil, &EmptyStringError{}
	}
	for idx, r := range e {
		if r == '@' {
			sub := e[:idx - 1]
			return &sub, nil
		}
	}
	return nil, &EmptyStringError{}
}