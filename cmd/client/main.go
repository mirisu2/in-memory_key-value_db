package main

import (
	"client-server-db/internal/client"
	"flag"
	"fmt"
	"os"
)

func main() {
	address := flag.String("address", "localhost:5555", "server address")
	flag.Parse()

	cl, err := client.NewClient(*address)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cl.Run()
}
