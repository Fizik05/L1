package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // Инициализируем context
	defer cancel()

	var wg sync.WaitGroup
	numbers := make(chan int) // Главный поток для записи данных
	var n int

	fmt.Print("Input count of workers: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ { // Запуск нужного количества воркеров
		wg.Add(1)
		go worker(ctx, i, numbers, &wg)
	}

	go func() {
		for i := 0; ; i += 3 {
			select {
			case <-ctx.Done(): // Если мы завершили программу, то выходим из цикла
				return
			case numbers <- i: // Запись данных в главный поток
				time.Sleep(1 * time.Second)
				break
			}
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	<-signalChannel
	cancel()  // Закрытие context
	wg.Wait() // Ожидание завершения работы всех воркеров
	fmt.Println("All workers finished")
}

func worker(ctx context.Context, id int, numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case number, ok := <-numbers:
			if !ok {
				return
			}
			fmt.Printf("Worker №%d received %d\n", id, number)
		}
	}
}
