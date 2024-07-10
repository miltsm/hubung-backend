package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/miltsm/hubung-service/internal/handler"
	"github.com/miltsm/hubung-service/internal/middleware"
)

const (
	GET_USERS="GET /v1/users"
	POST_USERS="POST /v1/users"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := http.NewServeMux()
	h := handler.New()
	
	r.HandleFunc(GET_USERS, h.RenderLogin)
	r.HandleFunc(POST_USERS, h.Login)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	s := http.Server {
		Addr: ":8080",
		Handler: stack(r),
	}

	fmt.Println("Server listening on port 8080")
	err = s.ListenAndServe()
	fmt.Println(err)
}