package main

import (
	"fmt"
	"math"
)

func main() {
	var key int
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Последовательность температурных колебаний
	storage := make(map[int][]float64)                                           // map, который будет хранить в себе ключи и подмножества

	for _, temp := range temperatures {
		if temp >= 0 {
			key = int(math.Floor(temp/10) * 10)
		} else {
			key = int(math.Ceil(temp/10) * 10)
		}
		if _, ok := storage[key]; !ok {
			storage[key] = make([]float64, 0)
		}
		storage[key] = append(storage[key], temp)
	}

	for key, value := range storage {
		fmt.Println(key, ":", value)
	}
}
