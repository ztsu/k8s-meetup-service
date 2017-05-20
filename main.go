package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"context"
	"os"
	"gopkg.in/op/go-logging.v1"
)

var (
	logger *logging.Logger
)

func init () {
	logging.SetBackend(
		logging.NewBackendFormatter(
			logging.NewLogBackend(os.Stderr, "", 0),
			logging.MustStringFormatter("%{color}[%{time:2006-01-02 15:04:05}] %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}"),
		),
	)

	logger = logging.MustGetLogger("")
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
		logger.Fatal("SERVICE_HOST_AND_PORT env var is empty")
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mv)
	r.Get("/", handler)

	logger.Info("Running at " + hostAndPort)

	http.ListenAndServe(hostAndPort, r)
}