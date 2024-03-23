package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик по окончании функции
	fmt.Println(n * n)
}

func main() {
	numbers := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup // объявляем WaitGroup для ожидания окончания всех Го рутин

	for _, number := range numbers {
		wg.Add(1)              // Увеличиваем счётчик
		go square(number, &wg) // Вызываем го рутину, передавая 2 параметром адрес переменной wg
		// Ведь передача по значению не пройдет, так как WaitGroup не должен копироваться после первого использования
	}

	wg.Wait() // Ожидаем завершения всех го рутин
}
