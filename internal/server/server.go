package server

import (
	"fmt"
	"hello-go/internal/request"
	"hello-go/internal/response"
	"io"
	"net"
)

type Server struct {
	closed  bool
	handler Handler
}

type HandlerError struct {
	StatusCode response.StatusCode
	Message    string
}

type Handler func(w *response.Writer, req *request.Request)

func runServer(s *Server, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		if s.closed {
			return
		}
		go runConnection(s, conn)
	}
}

func runConnection(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()
	responseWriter := response.NewWriter(conn)
	header := response.GetDefaultHeaders(0)
	req, err := request.RequestFromReader(conn)
	if err != nil {
		responseWriter.WriteStatusLine(response.StatusBadRequest)
		responseWriter.WriteHeaders(header)
		return
	}
	s.handler(responseWriter, req)
}

func Serve(port int, handler Handler) (*Server, error) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	server := &Server{closed: false, handler: handler}
	go runServer(server, listen)

	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
