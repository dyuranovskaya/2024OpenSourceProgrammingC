// func main() {
// 	answer := rand.Intn(6) + 1
// 	fmt.Println(answer)

// fmt.Print("number 입력 : ")
// in := bufio.NewReader(os.Stdin)
// i, err := in.ReadString('\n')
// if err != nil {
// 	log.Fatal(err)
// 	}

// 	i = strings.TrimSpace(i)
// 	quess, err := strconv.Atoi(i)

//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(quess)
//		if answer == quess {
//			fmt.Println("정답입니다!")
//		} else if answer > quess {
//			fmt.Println("입력하신 수는 정보다 작은 수 입니다. LOW")
//		} else {
//			fmt.Println("입력하신 수는 정보다 작은 수 입니다.HIGH")
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Println(answer)
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Println("d%번의 기회가 남았습니다. 숫자 입력:", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		i, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		i = strings.TrimSpace(i)

		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
		} else if answer > guess {
			fmt.Println("입력하신 수는 정보다 작은 수 입니다. LOW")
		} else {
			fmt.Println("입력하신 수는 정보다 작은 수 입니다. HIGH")
		}
	}

}
