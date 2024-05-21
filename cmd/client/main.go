package main

import (
	"client-server-db/internal/client"
	"client-server-db/internal/logger"
	"flag"
	"os"
)

func main() {
	address := flag.String("address", "localhost:5555", "server address")
	flag.Parse()

	cl, err := client.NewClient(*address, logger.Log)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	cl.Run()
}
