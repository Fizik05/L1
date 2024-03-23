package main

import "fmt"

func main() {
	a := 4
	b := 9

	fmt.Printf("The second option:\n\tOld values: a = %d, b = %d\n", a, b)

	a += b
	b = a - b
	a -= b

	fmt.Printf("\tNew values: a = %d, b = %d", a, b)
}
