package main

import (
	"fmt"
	"os"
)

func main() {

	var input string   //No termine la ejecucion
	fmt.Scanln(&input) //No termina la ejecucion

	defer fmt.Println("!")

	os.Exit(5)
}
