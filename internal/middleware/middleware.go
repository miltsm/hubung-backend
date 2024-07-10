package middleware

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i:=len(xs)-1;i>=0;i--{
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode: http.StatusOK,
		}
		next.ServeHTTP(wrapped,r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

// type middleware struct {
// 	handler http.Handler
// 	wauth *webauthn.WebAuthn
// }

// func New(base http.Handler) *middleware {
// 	wconfig := &webauthn.Config{
// 		RPDisplayName: "Go Webauthn",
// 		RPID: "hubung",
// 		RPOrigins: []string{  },
// 	} 

// 	wauth, err := webauthn.New(wconfig)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}

// 	return &middleware{
// 		handler: base,
// 		wauth: wauth,
// 	}
// }

