package main

import "fmt"

func binarySearch(arr []int, n int) (int, error) {
	var middle int
	left := 0 // Инициализируем границы
	right := len(arr) - 1

	for left <= right {
		middle = (left + right) / 2 // Находим средний элемент
		switch {
		case n < arr[middle]: // Если искомый элемент меньше среднего, то ищем в левой части
			right = middle - 1
		case n > arr[middle]: // Если искомый элемент большего среднего, то ищем в правой части
			left = middle + 1
		case n == arr[middle]: // Если равен возвращаем индекс элемента
			return middle, nil
		}
	}

	return -1, fmt.Errorf("Element is not found") // Если не нашли, возвращаем -1
}

func main() {
	arr := []int{3, 5, 7, 9, 11, 45, 46, 97}
	n := 100

	ind, err := binarySearch(arr, n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ind)
}
