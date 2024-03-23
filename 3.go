package main

import (
	"fmt"
	"sync"
)

func square2(n int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done() // По окончании функции уменьшаем счётчик
	ch <- n * n     // Отправляем в канал квадрат числа
}

func main() {
	numbers := [...]int{2, 4, 6, 8, 10} // Создаём массив с числами
	channel := make(chan int)           // Создаём канал, в который будем отправлять квадраты чисел
	sum := 0

	var wg sync.WaitGroup // Создаём WaitGroup для ожидания завершения всех го-рутин вычисления квадратов

	for _, number := range numbers {
		wg.Add(1)                        // Увеличиваем счётчик
		go square2(number, &wg, channel) // Вызываем го-рутину
	}

	go func() { // Создаём ещё одну го-рутину, которая будет ожидать завершения всех го-рутин вычисления квадратов и закрывать канал
		wg.Wait()
		close(channel)
	}() // Запускаем её

	for number := range channel { // Читаем данные из канала и вычисляем сумму до того момента, пока канал не закроется( за это отвечает последняя го-рутина)
		sum += number
	}
	fmt.Println(sum)
}
