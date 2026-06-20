package main

import (
	"crypto/sha256"
	"fmt"
	"hello-go/internal/headers"
	"hello-go/internal/request"
	"hello-go/internal/response"
	"hello-go/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const port = 42069

func toStr(data []byte) string {
	out := ""
	for _, b := range data {
		out += fmt.Sprintf("%02x", b)
	}
	return out
}

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
		} else if req.RequestLine.RequestTarget == "/video" {
			f, _ := os.ReadFile("assets/vim,mp4")
			h.Replace("content-type", "video/mp4")
			h.Replace("content-length", fmt.Sprintf("%d", len(f)))
			w.WriteStatusLine(response.StatusOk)
			w.WriteHeaders(h)
			w.WriteBody(f)
		} else if strings.HasPrefix(req.RequestLine.RequestTarget, "/httpbin/") {
			target := req.RequestLine.RequestTarget

			fmt.Printf("here %s\n", "https://httpbin.org/"+target[len("/httpbin/"):])
			res, err := http.Get("https://httpbin.org/" + target[len("/httpbin/"):])
			if err != nil {
				body = response.Response500()
				statusCode = response.StatusInternalServerError
			} else {
				h.Delete("Content-length")
				h.Set("Transfer-Encoding", "chunked")
				h.Replace("Content-Type", "text/plain")
				w.WriteStatusLine(response.StatusOk)
				w.WriteHeaders(h)
				h.Set("trailers", "X-Content-SHA256")
				h.Set("trailers", "X-Content-Length")
				fullBody := []byte{}
				// 開始讀資料
				for {
					data := make([]byte, 32)
					n, err := res.Body.Read(data)
					if err != nil {
						break
					}
					fullBody = append(fullBody, data[:n]...)
					w.WriteBody([]byte(fmt.Sprintf("%x\r\n", n)))
					w.WriteBody([]byte(data[:n]))
					w.WriteBody([]byte("\r\n"))
				}
				w.WriteBody([]byte("0\r\n"))
				tailers := headers.NewHeaders()
				out := sha256.Sum256(fullBody)
				tailers.Set("X-Content-SHA256", toStr(out[:]))
				tailers.Set("X-Content-Length", fmt.Sprintf("%d", len(fullBody)))
				w.WriteHeaders(tailers)
				w.WriteBody([]byte("\r\n"))
				return
			}
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
