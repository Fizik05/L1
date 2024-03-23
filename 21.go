package main

import (
	"fmt"
	"strconv"
)

type ExternalService interface { // Сторонний сервис, который работает с числами
	GetInteger() int
}

type Integers struct{}

func (i *Integers) GetInteger() int { // Возвращает число
	return 123
}

type OurService interface { // Наш сервис, который работает со строками
	GetString() string
}

type Strings struct{}

func (o *Strings) GetString() string { // Возвращает строку
	return "456"
}

type Adapter struct { // Адаптер
	externalService ExternalService
}

func (a *Adapter) GetString() string { // Метод для нашего сервиса
	data := a.externalService.GetInteger()

	return strconv.Itoa(data)
}

func main() {
	externalService := &Integers{}                        // Инициализация стороннего сервиса
	ourService := &Strings{}                              // Инициализация нашего сервиса
	adapter := &Adapter{externalService: externalService} // Инициализация адаптера

	fmt.Println(adapter.GetString() + ourService.GetString()) // Получение строки с нашего сервиса, используя сторонний сервис
}
