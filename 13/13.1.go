package main

import "fmt"

func main() {
	a := 4
	b := 9

	fmt.Printf("The first option:\n\tOld values: a = %d, b = %d\n", a, b)

	a, b = b, a

	fmt.Printf("\tNew values: a = %d, b = %d", a, b)
}
