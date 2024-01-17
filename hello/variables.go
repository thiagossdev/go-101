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
}
