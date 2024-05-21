package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

type Client struct {
	port   string
	logger *slog.Logger
}

func NewClient(port string, logger *slog.Logger) (*Client, error) {
	return &Client{
		port:   port,
		logger: logger,
	}, nil
}

func (s *Client) Run() {
	clientAddress := "localhost:5555"

	conn, err := net.Dial("tcp", clientAddress)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", clientAddress)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			break
		}

		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)
	}
}
