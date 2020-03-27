package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var stop uint

func main() {
	defer fmt.Println("!")
	conn, err := net.Dial("tcp", "192.168.50.169:9998")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = conn.Write([]byte("Hello Server !"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: Hello Server !")

	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		//HERE WE ARE RECEIVENG THE TIME
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))
		aux := string(data[:])           //Byte to string
		timeless, _ := strconv.Atoi(aux) //String to int
		fmt.Println(timeless)

		//HERE WE ARE RECEIVENG THE ID OF THE PROCESS
		buffer2 := make([]byte, 1400)
		dataSize, err = conn.Read(buffer2)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data2 := buffer2[:dataSize]
		fmt.Println("received message: ", string(data2))
		aux = string(data2[:])     //Byte to string
		id, _ := strconv.Atoi(aux) //String to int
		fmt.Println(id)

		//START OF THE PROCESS
		process := true
		const ValorMaximoInt = ^uint(0)
		//exitProcess := 0
		go func() {
			idProcess := id
			for i := uint(timeless); i < ValorMaximoInt; i++ {
				if process {
					fmt.Println("process: ", idProcess, "\tWith Time: ", i)
					stop = i
				}
				time.Sleep(time.Millisecond * 500)
			}
		}()

		//END OF THE PROCESS
		fmt.Scanln()
		s := strconv.Itoa(id)          //int to string
		_, err = conn.Write([]byte(s)) //takes a string
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", s)
		break
	}

}
