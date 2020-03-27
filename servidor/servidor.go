package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"time"
)

var process = true
var processid = 0
var process2 = true
var process3 = true
var process4 = true
var process5 = true
var tipo string
var NumberOfProcess = 0
var conexion net.Conn

func main() {
	listener, err := net.Listen("tcp", ":9998")
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
				fmt.Println("process: ", idProcess, "\tWith time: ", i)
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
				fmt.Println("process: ", idProcess, "\tWith time: ", i)
				timeless = i
			}
			if exitProcess == idProcess {
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		idProcess := 3
		for i := uint(0); i < ValorMaximoInt; i++ {
			if process3 {
				fmt.Println("process: ", idProcess, "\tWith time: ", i)
				timeless = i
			}
			if exitProcess == idProcess {
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		idProcess := 4
		for i := uint(0); i < ValorMaximoInt; i++ {
			if process4 {
				fmt.Println("process: ", idProcess, "\tWith time: ", i)
				timeless = i
			}
			if exitProcess == idProcess {
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		idProcess := 5
		for i := uint(0); i < ValorMaximoInt; i++ {
			if process5 {
				fmt.Println("process: ", idProcess, "\tWith time: ", i)
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
		//conexion = conn
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
		if NumberOfProcess == 3 {
			process3 = false
			processid = 3
		}
		if NumberOfProcess == 4 {
			process4 = false
			processid = 4
		}
		if NumberOfProcess == 5 {
			process5 = false
			processid = 5
			NumberOfProcess = 0
		}

		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		data := buffer[:dataSize]
		if err != nil {

			fmt.Println("connection closed")
			fmt.Println(err)
			return
		}
		fmt.Println("received message! ", string(data))
		aux := string(data[:])      //Byte to string
		id2, _ := strconv.Atoi(aux) //String to int
		fmt.Println(id2)
		fmt.Println(reflect.TypeOf(id2))
		if id2 == 1 {
			process = true
			return
		}
		if id2 == 2 {
			process2 = true
			return
		}
		if id2 == 3 {
			process3 = true
			return
		}
		if id2 == 4 {
			process4 = true
			return
		}
		if id2 == 5 {
			process5 = true
			return
		}

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
