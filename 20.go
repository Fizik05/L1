package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Split(str, " ") // Создаём срез из слов в исходной строке

	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i] // Меняем соответствующие слова
	}

	return strings.Join(words, " ") // Возвращем строку, созданную из среза
}

func main() {
	str := "snow dog sun"

	str = reverseWords(str)
	fmt.Println(str)
}
