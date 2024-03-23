package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

func main() {
	var wg sync.WaitGroup
	c := Counter{} // Инициализируем структуру, по дэфолту значение поля = 0

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&c.value, 1) // Атомарно инкрементируем
			wg.Done()
		}()
	}

	wg.Wait() // Дожидаемся завершения всех горутин
	fmt.Println(c.value)
}
