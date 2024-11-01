// package main

// import (
// 	"fmt"
// 	"log" //fmt — для вывода текста.
// 	//log — для отображения ошибок.
// )

// //Создаём функцию paintNeeded, которая принимает ширину и высоту и возвращает количество краски (число) и ошибку (если она есть).
// func paintNeeded(width float64, height float64) (float64, error) {
// //Если ширина меньше 0, возвращаем 0 и сообщение об ошибке с неверной шириной.
// 	if width < 0 {
// 		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
// 	}
// 	if height < 0 {
// 		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
// 	}
// 	//Если высота меньше 0, возвращаем 0 и сообщение об ошибке с неверной высотой.
// 	area := width * height
// 	return area / 10.0, nil //Вычисляем площадь, умножив ширину на высоту.
// }

// //Возвращаем количество краски (площадь делим на 10) и nil (нет ошибки).
//
//	func main() {
//		amount, err := paintNeeded(5.2, 3.5) //Запускаем paintNeeded с шириной 5.2 и высотой 3.5, и сохраняем результат в amount и ошибку в err.
//		if err != nil {
//			log.Fatal(err) //сли есть ошибка, программа остановится и покажет сообщение об ошибке.
//		}
//		fmt.Printf("%0.2f liters needed\n", amount)
//		amount, err = paintNeeded(4.2, -3.0)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("%0.2f liters needed\n", amount)
//	}
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// Создаём новый `Reader` для чтения ввода от пользователя
// 	in := bufio.NewReader(os.Stdin)
// 	fmt.Println("점수 입력:") // Сообщение для пользователя

// 	// Читаем строку ввода
// 	input, err := in.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err) // Завершаем программу, если произошла ошибка ввода
// 	}

// 	// Убираем пробелы и символ новой строки
// 	input = strings.TrimSpace(input)

// 	// Конвертируем строку в целое число
// 	n, err := strconv.Atoi(input)
// 	if err != nil {
// 		log.Fatal(err) // Завершаем программу, если произошла ошибка конвертации
// 	}

// 	// Считаем количество делителей числа `n`
// 	counts := 0
// 	j := 1
// 	for j <= n {
// 		if n%j == 0 {
// 			counts++
// 		}
// 		j++
// 	}

//		// Проверяем, является ли `n` простым числом
//		if counts == 2 {
//			fmt.Printf("%d is a prime number\n", n)
//		} else {
//			fmt.Printf("%d is not a prime number\n", n)
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	// Чтение ввода
// 	in := bufio.NewReader(os.Stdin)
// 	fmt.Println("Введите число:")

// 	input, err := in.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Убираем лишние пробелы и символ новой строки
// 	input = strings.TrimSpace(input)

// 	// Конвертируем строку в целое число
// 	n, err := strconv.Atoi(input)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Проверка на простое число
// 	isPrime := true
// 	for j := 2; j < n; j++ {
// 		if n%j == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	// Вывод результата
// 	if isPrime && n > 1 {
// 		fmt.Printf("%d is a prime number.\n", n)
// 	} else {
// 		fmt.Printf("%d is NOT a prime number.\n", n)
// 	}
// }

// func main() {
// 	// Чтение ввода от пользователя
// 	in := bufio.NewReader(os.Stdin)
// 	fmt.Println("Введите число:")

// 	input, err := in.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err) // Завершить программу при ошибке
// 	}

// 	// Убираем лишние пробелы и символ новой строки
// 	input = strings.TrimSpace(input)

// 	// Конвертируем строку в целое число
// 	n, err := strconv.Atoi(input)
// 	if err != nil {
// 		log.Fatal(err) // Завершить программу при ошибке
// 	}

// 	// Проверка на простое число
// 	isPrime := true
// 	for j := 2; j*j <= n; j++ {
// 		if n%j == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}

// 	// Вывод результата
// 	if isPrime && n > 1 {
// 		fmt.Printf("%d is a prime number.\n", n)
// 	} else {
// 		fmt.Printf("%d is NOT a prime number.\n", n)
// 	}
// }

func main() {
	in := bufio.NewReader(os.Stdin)
	// fmt.Printf("%f\n", math.Sqrt(19.0))
	fmt.Println("Input number:")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else if n == 0 {
		isPrime = true
	} else if n%2 == 0 {
		isPrime = false
	}

	if n <= 1 {
		isPrime = false
	} else {
		j := 3
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break
			}

			fmt.Printf("%d ", j)
			j = j + 2
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
