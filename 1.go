package main

import "fmt"

type Human struct {
	a, b int
}

type Action struct {
	Human // Встраивание методов структуры Human
}

func (h *Human) sum() {
	fmt.Printf("The summa of %d and %d is %d", h.a, h.b, h.a+h.b)
}

func main() {
	human := Human{10, 20}  // Создание экземпляра структуры Human
	action := Action{human} // Создание экземпляра структуры Action
	action.sum()            // Вызов метода структуры Human через экземпляр структры Action
}
