package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-flow-bot/internal/config"
)

func main() {
	/* CONFIG */
	cfg := config.NewConfig()
	slog.Info(
		"STARTING GO FLOW BOT...",
		slog.String("env", cfg.Env),
	)

	/* INIT ROUTER */
	r := chi.NewRouter()

	/* MIDDLEWARES */
	r.Use(middleware.Logger)

	/* ROUTES */
	r.Get("/", pingHandler)

	/* START SERVER */
	if err := http.ListenAndServe(cfg.HTTPServer.Address, r); err != nil {
		slog.Error("failed to start server")
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
