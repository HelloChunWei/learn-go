package server

import (
	"fmt"
	"hello-go/internal/response"
	"io"
	"net"
)

type Server struct {
	closed bool
}

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
	header := response.GetDefaultHeaders(0)
	response.WriteStatusLine(conn, response.StatusOk)
	response.WriteHeaders(conn, header)
}

func Serve(port int) (*Server, error) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	server := &Server{closed: false}
	go runServer(server, listen)

	return server, nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
