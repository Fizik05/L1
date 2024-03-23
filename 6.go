package main

import (
	"context"
	"fmt"
	"time"
)

func goFunc1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Остановка горутины с помощью context
			fmt.Println("First finished")
			return
		default:
			fmt.Println("first is working..")
			time.Sleep(1 * time.Second)
		}
	}
}

func goFunc2(stop <-chan bool) {
	for {
		select {
		case <-stop: // Остановка горутины с помощью дополнительного канала, отвечающего за остановку. Как только мы закинем туда значение, он разблокируется и завершит работу горутины
			fmt.Println("Second finished")
			return
		default:
			fmt.Println("second is working..")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go goFunc1(ctx)
	time.Sleep(1 * time.Second)
	cancel()

	stop := make(chan bool)
	go goFunc2(stop)
	time.Sleep(1 * time.Second)
	stop <- true

	time.Sleep(1 * time.Second)
}
