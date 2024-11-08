package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

// Функция для считывания целого числа с ввода, возвращает ошибку, если ввод некорректен
func getInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a1, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}
	a1 = strings.TrimSpace(a1)
	n, err := strconv.Atoi(a1)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	fmt.Print("Input start number: ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input end number: ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
