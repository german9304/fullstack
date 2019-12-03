package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	fullstack_backend "github.com/german9304/fullstack-backend"
)

const defaultPort = "8000"



func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	h := handler.GraphQL(fullstack_backend.NewExecutableSchema(fullstack_backend.Config{Resolvers: &fullstack_backend.Resolver{}}))

	handleCtx := func(k string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headers := w.Header()
			headers.Set("Access-Control-Allow-Origin", "http://localhost:3000")
			headers.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			headers.Set("Access-Control-Allow-Credentials", "true")
			ctx := r.Context()
			auth := fullstack_backend.Auth{w, r}
			newContext := context.WithValue(ctx, k, auth)
			cr := r.WithContext(newContext)
			h.ServeHTTP(w, cr)
		}
	}

	http.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", h)
	http.Handle("/", handleCtx("response"))

	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
