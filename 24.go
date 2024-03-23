package main

import (
	"fmt"
	"math"
)

type Point struct { // Создаём собственную структуру
	x, y float64
}

func NewPoint(x, y float64) *Point { // Конструктор для структуры
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	firstPoint := NewPoint(2, 8)   // Инициализация первой точки
	secondPoint := NewPoint(10, 3) // Инициализация второй точки

	distance := math.Sqrt(math.Pow(firstPoint.x-secondPoint.x, 2) + math.Pow(firstPoint.y-secondPoint.y, 2)) // Вычисление расстояния между ними

	fmt.Print(distance)
}
