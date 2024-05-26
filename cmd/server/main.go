package main

import (
	"client-server-db/internal/config"
	"client-server-db/internal/logger"
	"client-server-db/internal/server"
	"client-server-db/internal/storage"
	"flag"
	"fmt"
	"os"
)

func main() {
	file := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	logg, err := logger.NewLogger(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	logg.Info(fmt.Sprintf("Engine.Type: %s", cfg.Engine.Type))
	store, err := storage.NewStorage(cfg.Engine.Type, logg)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}

	srv, err := server.NewServer(cfg, store, logg)
	if err != nil {
		logg.Error(err.Error())
		os.Exit(1)
	}
	srv.Run()
}
