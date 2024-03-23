package main

import "fmt"

func main() {
	a := []int{2, 4, 1, 9, 4} // Инициализация среза
	var i int

	fmt.Println("Old slice:", a)

	fmt.Print("Input i: ")
	_, err := fmt.Scan(&i) // Инициализация индекса i из ввода пользователя
	if err != nil {
		fmt.Println("Integer!!!")
		return
	}
	if i < 0 || i >= len(a) {
		fmt.Println("Input right index!!!")
		return
	}

	a = append(a[:i], a[i+1:]...) // Удаление элемента из среза

	fmt.Println("New slice:", a)
}
