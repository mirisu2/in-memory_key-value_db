package server

import (
	"bufio"
	"client-server-db/internal/compute"
	"client-server-db/internal/config"
	"client-server-db/internal/storage"
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	address        string
	maxConnections int
	storage        storage.Storage
	logger         *slog.Logger
}

func NewServer(cfg *config.Config, storage storage.Storage, logg *slog.Logger) (*Server, error) {
	return &Server{
		address:        cfg.Network.Address,
		maxConnections: cfg.Network.MaxConnections,
		storage:        storage,
		logger:         logg,
	}, nil
}

func (s *Server) Run() {
	var addr string

	if s.address == "" {
		addr = "127.0.0.1:5555"
	} else {
		addr = s.address
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}
	defer listener.Close()

	s.logger.Info(fmt.Sprintf("listening on address %s", addr))

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error(err.Error())
			continue
		}
		s.logger.Info(fmt.Sprintf("new connection from %s", conn.RemoteAddr().String()))
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		req := scanner.Text()
		response := compute.HandlerMessages(req, s.storage)
		conn.Write([]byte(response + "\n"))
	}
}
