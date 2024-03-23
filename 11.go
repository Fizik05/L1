package main

import "fmt"

func intersaction(set1, set2 map[int]*struct{}) map[int]*struct{} {
	result := make(map[int]*struct{})
	for key1, _ := range set1 { // Ищем пересечения и заполняем новое множество
		for key2, _ := range set2 {
			if key1 == key2 {
				result[key1] = new(struct{})
			}
		}
	}

	return result
}

func main() {
	set1 := make(map[int]*struct{})
	set2 := make(map[int]*struct{})

	for i := 0; i < 20; i += 1 { // Заполнение первого множества
		set1[i] = new(struct{})
	}

	for i := 5; i < 20; i += 2 {
		set2[i] = new(struct{}) // Заполнение второго множества
	}

	commonSet := intersaction(set1, set2)
	for i, _ := range commonSet {
		fmt.Println(i)
	}
}
