package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"context"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if id, ok := r.Context().Value("id").(string); ok {
		w.Write([]byte("Hello, " + id + "\n"))
		return
	} else {
		w.Write([]byte("Hello guest\n"))
	}
}

func mv(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "id", "Qq")
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mv)
	r.Get("/", handler)

	http.ListenAndServe(":8080", r)
}