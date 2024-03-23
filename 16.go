package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high) // Функция которая меняет элементы, относительно опорного элемента и потом возвращает его
		quickSort(arr, low, pi-1)       // Запускаем рекурсию для подмассивов, которые больше и мменьше опроного элемента
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Опорный элемент - последний
	i := low - 1
	for j := low; j <= high-1; j++ { // Проходимся по выделенной части массива
		if arr[j] < pivot { // Элементы, меньше опроного ставим в начало, а элементы, которые больше опорного попадают в конец массива
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Установили опорный элемент в середину, между меньшими и большими
	return i + 1
}

func main() {
	arr := []int{3, 6, 1, 9, 10, 45, 0, 0, 0, 0}

	fmt.Println("Before sort:", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println("After sort:", arr)
}
