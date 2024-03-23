package main

import (
	"fmt"
	"strings"
)

func isSet(str string) bool {
	set := make(map[rune]bool) // Инициализирую множество

	for _, i := range str {
		if _, ok := set[i]; !ok {
			set[i] = true // Если символ раньше не встречался, добавляем его в множество
		} else {
			return false // Иначе возвращаем false
		}
	}

	return true
}

func main() {
	var str string

	fmt.Print("Input a string: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Print("Just a string!!!")
		return
	}

	fmt.Printf("%s - %t", str, isSet(strings.ToLower(str)))
}
