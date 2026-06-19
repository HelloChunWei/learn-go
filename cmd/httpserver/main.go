package main

import (
	"fmt"
	"hello-go/internal/request"
	"hello-go/internal/response"
	"hello-go/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = 42069

func main() {
	server, err := server.Serve(port, func(w *response.Writer, req *request.Request) {
		h := response.GetDefaultHeaders(0)
		body := response.Response200()
		statusCode := response.StatusOk
		if req.RequestLine.RequestTarget == "/yourproblem" {
			body = response.Response400()
			statusCode = response.StatusBadRequest
		} else if req.RequestLine.RequestTarget == "/myproblem" {
			body = response.Response500()
			statusCode = response.StatusInternalServerError
		}
		h.Replace("Content-Type", "text/html")
		h.Replace("Content-length", fmt.Sprintf("%d", len(body)))
		w.WriteStatusLine(statusCode)
		w.WriteHeaders(h)
		w.WriteBody(body)
	})
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
