package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Print("Input value of n: ")
	fmt.Scanln(&n)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup
	currentTime := time.Now()

	wg.Add(2)
	go read(ctx, ch, &wg)
	go write(ctx, ch, &wg)

	wg.Wait()
	fmt.Print("All done\n")
	fmt.Print("Time: ", time.Since(currentTime))
}

func read(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

func write(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
}
