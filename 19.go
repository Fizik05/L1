package main

import "fmt"

func reverse(str string) string {
	runes := []rune(str) // Инициализируем срез рун по исходной строке
	left, right := 0, len(str)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left] // Меняем местами соответствующие элементы
		left++
		right--
	}

	return string(runes)
}

func main() {
	var str string

	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Input string")
	}

	str = reverse(str)
	fmt.Print(str)
}
