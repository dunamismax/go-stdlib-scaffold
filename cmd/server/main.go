package main

import (
	"context"
	"embed"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dunamismax/go-stdlib-scaffold/internal/database"
)

var (
	//go:embed templates/*
	templatesFS embed.FS
)

func main() {
	// Setup logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup database
	if err := database.Connect(); err != nil {
		logger.Error("failed to setup database", "error", err)
		os.Exit(1)
	}

	// Load templates
	tmpl, err := template.ParseFS(templatesFS, "templates/*")
	if err != nil {
		logger.Error("failed to parse templates", "error", err)
		os.Exit(1)
	}

	// Setup router
	mux := http.NewServeMux()

	// Static files
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets"))))

	// Routes
	mux.HandleFunc("/", handleIndex(tmpl))
	mux.HandleFunc("/messages", handleCreateMessage(tmpl))

	// Start server
	port := os.Getenv("APP_SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: loggingMiddleware(mux, logger),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("listen and serve error", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	logger.Info("server exited properly")
}

func loggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}

func handleIndex(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		messages, err := database.GetMessages()
		if err != nil {
			http.Error(w, "failed to get messages", http.StatusInternalServerError)
			return
		}
		if err := tmpl.ExecuteTemplate(w, "index.html", messages); err != nil {
			http.Error(w, "failed to execute template", http.StatusInternalServerError)
		}
	}
}

func handleCreateMessage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		content := r.FormValue("content")
		if content == "" {
			http.Error(w, "content is required", http.StatusBadRequest)
			return
		}

		if _, err := database.CreateMessage(content); err != nil {
			http.Error(w, "failed to create message", http.StatusInternalServerError)
			return
		}

		messages, err := database.GetMessages()
		if err != nil {
			http.Error(w, "failed to get messages", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "messages.html", messages); err != nil {
			http.Error(w, "failed to execute template", http.StatusInternalServerError)
		}
	}
}
