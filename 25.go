package main

import (
	"fmt"
	"time"
)

func sleep(timeDuration time.Duration) {
	<-time.After(timeDuration) // Функция блокируется на время, переданное в параметрах
}

func main() {
	fmt.Println("Starting...")
	start := time.Now()

	sleep(5 * time.Second)

	fmt.Println("Program worked", time.Since(start))
	fmt.Println("Done")
}
