package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Функция для проверки простого числа
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

// Функция для считывания целого числа
func getInteger(prompt string) int {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	// Ввод начального и конечного числа
	n1 := getInteger("Input start number: ")
	n2 := getInteger("Input end number: ")

	// Вывод простых чисел в указанном диапазоне
	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
	fmt.Println()
}
