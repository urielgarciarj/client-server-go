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
var Process = []int{1, 2, 3, 4, 5}

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
	for i := 0; i < 4; i++ {
		if Process[i] != 0 {
			switch Process[i] {
			case 1:
				fmt.Println("Si entre")
				processid = 1
				process = false
				Process[i] = 0
				break
			case 2:
				fmt.Println("Si entre")
				processid = 2
				process2 = false
				Process[i] = 0
				break
			case 3:
				fmt.Println("Si entre")
				processid = 3
				process3 = false
				Process[i] = 0
				break
			case 4:
				fmt.Println("Si entre")
				processid = 4
				process4 = false
				Process[i] = 0
				break
			case 5:
				fmt.Println("Si entre")
				processid = 5
				process5 = false
				Process[i] = 0
				break
			}
			break
		}
	}
	for {
		fmt.Println(processid)
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
		switch id2 {
		case 1:
			process = true
			Process[0] = 1
			return
		case 2:
			process2 = true
			Process[1] = 2
			return
		case 3:
			process3 = true
			Process[2] = 3
			return
		case 4:
			process4 = true
			Process[3] = 4
			return
		case 5:
			process5 = true
			Process[4] = 5
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
