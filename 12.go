package main

import "fmt"

type Set map[string]*struct{} // Создали свое собственное множество

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	result := Set{}

	for _, word := range strings { // Заполняем множество
		result[word] = new(struct{})
	}

	for key, _ := range result { // Выводим результат
		fmt.Println(key)
	}
}
