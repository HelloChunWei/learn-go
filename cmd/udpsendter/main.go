package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udp, err := net.ResolveUDPAddr("udp", "localhost:42069")

	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, udp)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Message sent: %s", message)
	}
}
