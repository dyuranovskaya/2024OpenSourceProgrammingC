package main

import (
	"fmt" // Подключаем пакет для форматированного вывода
	//"reflect" // Подключаем пакет для работы с типами данных
)

func main() {
	var f float64
	var i int
	var b bool
	var s string
	//mySchoolAccount := 5
	//println(mySchoolAccount)
	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i)
	f = 2.7
	i = 3
	fmt.Print("\n\n", (f <= float64(i)), "\n")
	//i := 13
	//var f float64 = 12.9 // Объявляем переменную с плавающей запятой 'f' и присваиваем значение 12.9
	//c1 := 'z'
	//c2 := "kim" //44688
	// Корректный форматированный вывод с помощью fmt.Printf
	//fmt.Printf("value i: %d, value f: %f\n", i, f)

	// Приведение типа 'f' к int для умножения и корректный вывод
	//fmt.Printf("%d * %f = %d\n", i, f, 1*int(f))
	//fmt.Println(c1, c2)
	// Вывод типа переменных
	//fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2)) // Выводим тип переменных 'i' и 'f'
}
