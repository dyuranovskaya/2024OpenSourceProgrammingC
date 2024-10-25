package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := rand.Intn(6) + 1
	fmt.Println(answer)

	fmt.Print("number 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	quess, err := strconv.Atoi(i)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quess)
	if answer == quess {
		fmt.Println("정답입니다!")
	} else if answer > quess {
		fmt.Println("입력하신 수는 정보다 작은 수 입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정보다 작은 수 입니다.HIGH")
	}
}
