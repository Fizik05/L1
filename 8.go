package main

import "fmt"

func main() {
	var n int64
	var i uint

	fmt.Print("Input an integer: ")
	fmt.Scanln(&n) // Инициализируем сило, которое хотим изменить

	fmt.Print("Input a number of bit: ")
	fmt.Scanln(&i) // инициализируем номер бита, отсчёт начинается справа с 1

	fmt.Printf("%d(%064b) ", n, n) // Выводим старое значение

	n ^= (1 << (i - 1)) // Инвертируем нужный нам бит, 0 меняем на 1 или 1 меняем на 0

	fmt.Printf("converted to %d(%064b)", n, n) // Выводим результат
}
