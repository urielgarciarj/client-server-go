package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"time"
)

var process = true
var processid = 0
var process2 = true
var tipo string
var NumberOfProcess = 0

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
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
	go func() {
		idProcess := 2
		for i := uint(0); i < ValorMaximoInt; i++ {
			if process2 {
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
		tipo := listener.Addr()

		fmt.Println(tipo)
		fmt.Println(reflect.TypeOf(tipo))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connection")
		go listenConnection(conn, timeless)
	}
}

func listenConnection(conn net.Conn, _timeless uint) {

	NumberOfProcess = NumberOfProcess + 1
	for {
		if NumberOfProcess == 1 {
			process = false
			processid = 1
		}
		if NumberOfProcess == 2 {
			process2 = false
			processid = 2
		}
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			fmt.Println(err)
			return
		}
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))

		time := fmt.Sprint(_timeless) //Convert a uint type into a string
		_, err = conn.Write([]byte(time))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", string(time))

		id := fmt.Sprint(processid) //Convert a uint type into a string
		_, err = conn.Write([]byte(id))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", string(id))
	}
}
