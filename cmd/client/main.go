package main

import (
	"client-server-db/internal/client"
	"client-server-db/internal/logger"
	"os"
)

func main() {

	// TODO Use environment variables or config args
	cl, err := client.NewClient("localhost:5555", logger.Log)
	if err != nil {
		logger.Log.Error(err.Error())
		os.Exit(1)
	}
	cl.Run()
}
