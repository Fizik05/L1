package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	mainMap := make(map[string]int)

	writeToMap := func(key string, value int) { // Функция для записи в map
		mutex.Lock()         //Блокируем Мьютекс
		defer mutex.Unlock() // В конце разблокируем

		mainMap[key] = value // Пишем данные в map
		fmt.Printf("There is a %d wtih key %s\n", value, key)

		wg.Done() // У меньшаем счётчик горутин
	}

	for i := 0; i < 10; i++ { // Запускаем горутины со счётчиком
		wg.Add(1)
		go writeToMap(strconv.Itoa(i), i*i)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All done")
}
