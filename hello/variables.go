package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Thiago Sousa Santos"
	version := 0.1
	fmt.Println("Hello,", name)
	fmt.Println("This program is in version", version)
	fmt.Println("The type of name var is", reflect.TypeOf(name))

	fmt.Println("1- Start monitor")
	fmt.Println("2- Display logs")
	fmt.Println("0- Exit the program")

	var command int
	fmt.Scanf("%d", &command)

	fmt.Println("\nThe chose command was", command)
}
