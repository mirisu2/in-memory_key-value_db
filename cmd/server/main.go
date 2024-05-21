package main

import (
	"client-server-db/internal/config"
	"client-server-db/internal/logger"
	"client-server-db/internal/server"
	"client-server-db/internal/storage"
	"flag"
	"os"
)

func main() {
	file := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.NewConfig(*file)
	if err != nil {
		logger.Log.Error(err.Error())
		//fmt.Println(err.Error())
		os.Exit(1)
	}

	store := storage.NewMemoryStorage()

	//logger, err := logger.NewLogger(cfg.Logging.Level, cfg.Logging.Format, cfg.Logging.Output)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}

	srv, err := server.NewServer(cfg, store, logger.Log)
	if err != nil {
		logger.Log.Error(err.Error())
		//logger.Log.Error(err.Error())
		os.Exit(1)
	}
	srv.Run()
}
