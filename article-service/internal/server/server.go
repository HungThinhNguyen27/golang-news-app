package server

import (
	"article-service/internal/configs"
	"article-service/internal/routes"
	"article-service/internal/services"
	"article-service/internal/storage/postgres"
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App struct to manage dependencies
type App struct {
	Config         *configs.Config
	Storage        *postgres.Postgres
	ArticleService *services.ArticleService
	Router         *http.ServeMux
	Server         *http.Server
}

// NewApp initializes all dependencies
func NewApp() (*App, error) {
	cfg := configs.MustLoad()

	// Initialize database
	storage, err := postgres.ConnectToDB()
	if err != nil {
		return nil, err
	}

	// Initialize service
	articleService := services.NewArticleService(storage)

	// Setup router
	router := routes.SetupRouter(articleService)

	// Create server
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	return &App{
		Config:         cfg,
		Storage:        storage,
		ArticleService: articleService,
		Router:         router,
		Server:         server,
	}, nil
}

// StartServer runs the HTTP server and handles graceful shutdown
func (app *App) StartServer() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		slog.Info("Server started", slog.String("address", app.Config.Addr))
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	// Wait for shutdown signal
	<-done
	slog.Info("Shutting down the server")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	} else {
		slog.Info("Server shutdown successfully")
	}
}
