package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 周期的把当前时间写给客户端
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatalf("err create listener %v", listener)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		handleConnection(connection)

	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	for {
		_, err := io.WriteString(connection, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(time.Second)
	}
}
