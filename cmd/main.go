package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/miltsm/hubung-service/internal/handler"
	"github.com/miltsm/hubung-service/internal/middleware"
)

const (
	PORT="SERVER_PORT"
	GET_SESSION="GET /sessions"	//render login page
	POST_SESSION="POST /sessions" //request server challenge
	PUT_SESSION="PUT /sessions" //send to server to validate key
)

func main() {
	r := http.NewServeMux()
	h := handler.New()
	
	r.HandleFunc(GET_SESSION, h.RenderLogin)
	r.HandleFunc(POST_SESSION, h.Login)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", r))

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	port := os.Getenv(PORT)

	s := http.Server {
		Addr: fmt.Sprintf(":%s", port),
		Handler: stack(v1),
	}

	fmt.Printf("Server listening on port %s\n", port)
	err := s.ListenAndServe()
	fmt.Println(err)
}