package main

import (
	"client-server-db/internal/logger"
	"client-server-db/internal/server"
	"client-server-db/internal/storage"
	"os"
)

func main() {
	store := storage.NewMemoryStorage()

	// TODO Use environment variables or config args
	srv, err := server.NewServer("5555", store, logger.Log)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	srv.Run()
}
