package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var process = true

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	const ValorMaximoInt = ^uint(0)
	exitProcess := 0
	var timeless uint
	go func() {
		idProcess := 1
		for i := uint(0); i < ValorMaximoInt; i++ {
			if process {
				fmt.Println("process: ", idProcess, "\tCon tiempo: ", i)
				timeless = i
			}
			if exitProcess == idProcess {
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connection")
		go listenConnection(conn, timeless)
	}
}

func listenConnection(conn net.Conn, _timeless uint) {
	for {
		process = false
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			process = true
			return
		}
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))

		s := fmt.Sprint(_timeless) //Convert a uint type into a string

		_, err = conn.Write([]byte(s))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", string(s))
	}
}
