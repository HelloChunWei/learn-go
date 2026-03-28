package main

import (
	"bytes"
	"fmt"
	"hello-go/internal/request"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(ch)

		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				ch <- str
				str = ""
			}
			str += string(data)
		}

		if len(str) != 0 {
			ch <- str
		}

	}()
	return ch
}

func main() {
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		connect, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		requestData, err := request.RequestFromReader(connect)
		if err != nil {
			log.Printf("parse request: %v", err)
			connect.Close()
			continue
		}
		fmt.Printf("Request line:\n")
		fmt.Printf("- Method: %s\n", requestData.RequestLine.Method)
		fmt.Printf("- Target: %s\n", requestData.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", requestData.RequestLine.HttpVersion)
	}
	// filePath := "messages.txt"
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatal("Error: ")
	// }

	// lines := getLinesChannel(file)

	// for line := range lines {
	// 	fmt.Printf("read: %s\n", line)
	// }
}
