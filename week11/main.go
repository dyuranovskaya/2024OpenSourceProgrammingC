package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Функция для проверки, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			j = j + 2
		}
	}
	return true
}

// Функция для считывания целого числа с ввода
func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n') // Чтение строки
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)  // Убираем лишние пробелы
	n, err := strconv.Atoi(a) // Преобразуем строку в число
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	fmt.Print("Введите начальное число: ")
	n1 := getInteger() // Получаем начальное число

	fmt.Print("Введите конечное число: ")
	n2 := getInteger() // Получаем конечное число

	// Печатаем все простые числа в диапазоне от n1 до n2
	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
