package server

import (
	"bufio"
	"client-server-db/internal/compute"
	"client-server-db/internal/logger"
	"client-server-db/internal/storage"
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	port    string
	storage storage.Storage
	logger  *slog.Logger
}

func NewServer(port string, storage storage.Storage, logger *slog.Logger) (*Server, error) {
	return &Server{
		port:    port,
		storage: storage,
		logger:  logger,
	}, nil
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	defer listener.Close()

	logger.Log.Info(fmt.Sprintf("listening on port %s", s.port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Log.Error(err.Error())
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		req := scanner.Text()
		response := s.handlerMessages(req)
		conn.Write([]byte(response + "\n"))
	}
}

func (s *Server) handlerMessages(req string) string {
	command, args, err := compute.Parse(req)
	if err != nil {
		return err.Error()
	}

	query, err := compute.Analyze(command, args, s.storage)
	if err != nil {
		return err.Error()
	}

	return query
}
