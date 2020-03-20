package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	defer fmt.Println("!")
	conn, err := net.Dial("tcp", ":9999")
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
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))
		aux := string(data[:])           //Byte to string
		timeless, _ := strconv.Atoi(aux) //String to int

		fmt.Println(timeless)

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

		process := true
		const ValorMaximoInt = ^uint(0)
		exitProcess := 0
		go func() {
			idProcess := id
			for i := uint(timeless); i < ValorMaximoInt; i++ {
				if process {
					fmt.Println("process: ", idProcess, "\tCon tiempo: ", i)
				}
				if exitProcess == idProcess {
					return
				}
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}
	os.Exit(5)
}
