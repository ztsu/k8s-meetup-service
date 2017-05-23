package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"context"
	"os"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func init () {
	//

	log.SetHandler(cli.New(os.Stderr))
}

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
	hostAndPort := os.Getenv("SERVICE_HOST_AND_PORT")
	if hostAndPort == "" {
		log.Fatal("SERVICE_HOST_AND_PORT env var is empty")
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mv)
	r.Get("/", handler)

	log.Infof("Running at %s", hostAndPort)

	err := http.ListenAndServe(hostAndPort, r)
	if err != nil {
		log.WithError(err).Fatal("Can not start service")
	}
}