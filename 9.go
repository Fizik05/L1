package main

import (
	"fmt"
)

func creating(in chan<- int) { // Фунция записи чисел в канал
	var a = []int{2, 4, 10, 0, 5, 8, 100}
	for _, v := range a {
		in <- v // Отправляем данные в канал
	}
	close(in) // По завершении работы функции закрываем канал
}

func handling(out <-chan int, in chan<- int) {
	for v := range out { // Читаем данные из первого канала
		in <- v * 2 // Отправляем во второй канал умноженные числа
	}
	close(in) // По завершении работы закрываем канал с умноженными числами
}

func reading(out <-chan int) {
	for v := range out { //  Читаем данные из второго канала и выводим их на экран
		fmt.Printf("%d ", v)
	}
}

func main() {
	numbers := make(chan int)           // Канал для исходных данных
	multipliedNumbers := make(chan int) // Канал для результирующих данных

	go creating(numbers)                    // Запускаем горутину, запись в первый канал
	go handling(numbers, multipliedNumbers) // Запускаем горутину, обработку данных из первого канала
	reading(multipliedNumbers)              // Запускаем функцию, которая распечатает результат на экран + обеспечивает завершение всех горутин
}
