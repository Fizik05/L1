package main

import (
	"fmt"
	"math/big"
)

func summa(a, b *big.Int) (c *big.Int) { // Сложение
	c = new(big.Int)
	return c.Add(a, b)
}

func subtraction(a, b *big.Int) (c *big.Int) { // Вычитание
	c = new(big.Int)
	return c.Sub(a, b)
}

func multiplication(a, b *big.Int) (c *big.Int) { // Умножение
	c = new(big.Int)
	return c.Mul(a, b)
}

func division(a, b *big.Int) (c *big.Int) { // Деление
	c = new(big.Int)
	return c.Div(a, b)
}

func main() {
	var a = new(big.Int) // Инициализация переменных типа math/big
	var b = new(big.Int)

	a.SetInt64(2 << 45)
	b.SetInt64(4 << 36)

	fmt.Println(summa(a, b))
	fmt.Println(subtraction(a, b))
	fmt.Println(multiplication(a, b))
	fmt.Println(division(a, b))
}
