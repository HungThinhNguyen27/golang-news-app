package server

import (
	"article-service/internal/repo/elasticsearch"
	"article-service/internal/routes"
	"article-service/internal/services"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App struct to manage dependencies
type App struct {
	ArticleService *services.ArticleServiceWithES
	Router         http.Handler
	Server         *http.Server
}

// NewApp initializes all dependencies
func NewApp(addr string) (*App, error) {

	// ES
	esStorage, err := elasticsearch.InitElasticsearch()
	if err != nil {
		fmt.Println("Error initializing Elasticsearch:", err)
		return nil, err // Đừng quên return nếu có lỗi
	}

	// Initialize service
	articleService := services.NewArticleServiceWithES(esStorage)
	router := routes.SetupRouter(articleService)

	return &App{
		ArticleService: articleService,
		Router:         router,
		Server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}, nil
}

// StartServer runs the HTTP server and handles graceful shutdown
func (app *App) StartServer(addr string) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		slog.Info("Server started", slog.String("address", addr))
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
