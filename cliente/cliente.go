package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = conn.Write([]byte("Hello Server !"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: Hello Server !")

	_, err = conn.Write([]byte("How are you?"))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Message sent: How are you?")

	/*var i = 9

	s := strconv.Itoa(i)

	fmt.Println(reflect.TypeOf(s))

	fmt.Println(s)*/

	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))

		aux := string(data[:]) //Byte to string

		timeless, _ := strconv.Atoi(aux) //String to int

		fmt.Println(timeless)

		process := true
		const ValorMaximoInt = ^uint(0)
		exitProcess := 0
		go func() {
			idProcess := 1
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
}
