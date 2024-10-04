package main

import (
	"fmt"     // Подключаем пакет для форматированного вывода
	"reflect" // Подключаем пакет для работы с типами данных
)

func main() {
	i := 13
	var f float64 = 12.9 // Объявляем переменную целого типа 'i' и присваиваем ей значение 55
	// Объявляем переменную с плавающей запятой 'f' и присваиваем значение 12.9
	fmt.Println("value i :%d,value f:%f\n", i, f) // Попытка форматированного вывода (ошибка)
	//fmt.Println("%d*%f=%f/n", i, f, i*f)
	//fmt.Println("%d*%f=%f/n", i, f, float64(i)*f) // Приводим i к типу float64 для корректного умножения с f
	fmt.Println("%d*%f=%d/n", i, f, 1*int(f))
	fmt.Println(reflect.TypeOf(i), (reflect.TypeOf(f)))
}

// Определяем и выводим тип переменной 'f' (ожидается float64)
