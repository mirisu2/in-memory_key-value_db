package server

import (
	"bufio"
	"client-server-db/internal/compute"
	"client-server-db/internal/config"
	"client-server-db/internal/logger"
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

func NewServer(cfg *config.Config, storage storage.Storage, logger *slog.Logger) (*Server, error) {
	return &Server{
		address:        cfg.Network.Address,
		maxConnections: cfg.Network.MaxConnections,
		storage:        storage,
		logger:         logger,
	}, nil
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	defer listener.Close()

	logger.Log.Info(fmt.Sprintf("listening on address %s", s.address))

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Error(err.Error())
			continue
		}
		logger.Log.Info(fmt.Sprintf("new connection from %s", conn.RemoteAddr().String()))
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
