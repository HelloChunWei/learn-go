package server

import (
	"bytes"
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

type Handler func(w io.Writer, req *request.Request) *HandlerError

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
	req, err := request.RequestFromReader(conn)
	if err != nil {
		response.WriteStatusLine(conn, response.StatusBadRequest)
		response.WriteHeaders(conn, header)
		return
	}
	writer := bytes.NewBuffer([]byte{})
	handlerError := s.handler(writer, req)
	var body []byte = nil
	statusCode := response.StatusOk
	if handlerError != nil {
		statusCode = handlerError.StatusCode
		body = []byte(handlerError.Message)
	} else {
		body = writer.Bytes()
	}
	header.Replace("Content-length", fmt.Sprintf("%d", len(body)))
	response.WriteStatusLine(conn, statusCode)
	response.WriteHeaders(conn, header)
	conn.Write(body)
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
