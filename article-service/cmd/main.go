package main

import (
	"article-service/internal/server"
	"log/slog"
	"os"
)

func main() {
	addr := "0.0.0.0:9090"
	app, err := server.NewApp(addr)
	if err != nil {
		slog.Error("Application failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	}
	app.StartServer(addr)
}
