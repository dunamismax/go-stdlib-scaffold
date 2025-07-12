package main

import (
	"context"
	"embed"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dunamismax/go-stdlib-scaffold/internal/database"
)

//go:embed templates/*.html
var templatesFS embed.FS

// application holds the application's dependencies.
type application struct {
	logger *log.Logger
	store  *database.Store
	tmpl   *template.Template
}

func main() {
	// Configuration
	port := os.Getenv("APP_SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	dbPath := os.Getenv("APP_DB_PATH")
	if dbPath == "" {
		dbPath = "data/app.db"
	}

	// Dependencies
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.NewDB(dbPath)
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	store := database.NewStore(db)

	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		logger.Fatalf("failed to parse templates: %v", err)
	}

	app := &application{
		logger: logger,
		store:  store,
		tmpl:   tmpl,
	}

	// Start server
	if err := app.serve(port); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}

// serve starts and gracefully shuts down the HTTP server.
func (app *application) serve(port string) error {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		app.logger.Println("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownErr <- srv.Shutdown(ctx)
	}()

	app.logger.Printf("starting server on port %s", port)
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	app.logger.Println("server stopped")
	return nil
}

// routes registers all application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Serve static files from the 'public' directory using the modern os.DirFS.
	mux.Handle("/css/", http.FileServer(http.FS(os.DirFS("public"))))

	// Register application handlers.
	mux.HandleFunc("GET /", app.handleIndex)
	mux.HandleFunc("POST /messages", app.handleCreateMessage)

	return app.loggingMiddleware(mux)
}

// loggingMiddleware logs incoming HTTP requests.
func (app *application) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		app.logger.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// handleIndex renders the main page with all messages.
func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {
	messages, err := app.store.GetMessages()
	if err != nil {
		app.logger.Printf("error getting messages: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := app.tmpl.ExecuteTemplate(w, "index.html", messages); err != nil {
		app.logger.Printf("error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// handleCreateMessage handles the creation of a new message.
func (app *application) handleCreateMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "Content cannot be empty", http.StatusBadRequest)
		return
	}

	if _, err := app.store.CreateMessage(content); err != nil {
		app.logger.Printf("error creating message: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
