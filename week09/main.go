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
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// rand.Seed(time.Now().Unix())
	// answer := rand.Intn(6) + 1
	// // fmt.Println(answer)

	// var win bool = false
	// for guesses := 0; guesses < 3; guesses++ {
	// 	fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력:", 3-guesses)
	// 	in := bufio.NewReader(os.Stdin)
	// 	i, err := in.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	i = strings.TrimSpace(i)

	// 	guess, err := strconv.Atoi(i)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(guess)

	// 	if answer == guess {
	// 		fmt.Println("정답입니다!")
	// 		win = true
	// 		break
	// 	} else if answer > guess {
	// 		fmt.Println("입력하신 수는 정보다 작은 수 입니다. LOW")
	// 	} else {
	// 		fmt.Println("입력하신 수는 정보다 작은 수 입니다. HIGH")
	// 	}
	// 	if win {
	// 		fmt.Println("당신이 이겼습니다!")
	// 	} else {
	// 		fmt.Printf("당신이 졌습니다. 정담이 %d입니다\n", answer)
	// 	}
	// }
	result := fmt.Sprintf("%5.2f나누기 %3.2f은(는) %3.2f입니다 \n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)

	i := 1
	// fmt.Printf("%T\n", i)
	for i <= 10 {
		fmt.Printf("%4d\n", rand.Intn(1000)+1)
		i++
	}
	fmt.Println(reflect.TypeOf(i))
}
