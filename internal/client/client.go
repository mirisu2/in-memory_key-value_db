package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
)

type Client struct {
	address string
	logger  *slog.Logger
}

func NewClient(address string, logger *slog.Logger) (*Client, error) {
	return &Client{
		address: address,
		logger:  logger,
	}, nil
}

func (s *Client) Run() {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at", s.address)

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
