package main

import (
	"article-service/internal/server"
	"log/slog"
	"os"
)

func main() {
	app, err := server.NewApp()
	if err != nil {
		slog.Error("Application failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
	app.StartServer()
}
