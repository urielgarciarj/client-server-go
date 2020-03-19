package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connection")

		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))

		_, err = conn.Write([]byte("I'm fine!"))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: I'm fine!")
	}
}
