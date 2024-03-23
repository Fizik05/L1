package main

import (
	"fmt"
	"reflect"
)

func varType(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("channel")
	}
}

func main() {
	var x interface{}

	x = 5
	varType(x)

	x = "Hello"
	varType(x)

	x = true
	varType(x)

	x = make(chan int)
	varType(x)
}
