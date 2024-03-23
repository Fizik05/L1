package main

func createHugeString(n int) string {
	// ...
	return "anything string" // Return a huge string
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // Первая проблема в том, что функция может вернуть строку длиной < 100 симмволов, тогда упадёт ошибка
	justString = v[:100]           // Вторая проблема: justString и v ссылаются на один и тот же массив байтов
	// Соответственно, в памяти будет хранится слишком длиная строка, из которой будет использоваться только первые 100 символов
}

func someRightFunc() {
	v := createHugeString(1 << 10)
	if len(v) < 100 { // Если длина строки < 100 элементов, то оставляем тот же способ
		justString = v
		return
	}
	buffer := make([]byte, 100) // Создаём срез байтов размерностью 100
	copy(buffer, v[:100])       // Копируем первые 100 элементов в буфер
	justString = string(buffer) // Теперь justString ссылается на массив байтов размерностью 100
}

func main() {
	someRightFunc()
}
