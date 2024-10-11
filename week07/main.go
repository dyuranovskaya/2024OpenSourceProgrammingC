package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var now time.Time = time.Now()
	//fmt.Printf("today %dyear %dmonth %dday.\n", now.Year(), now.Month(), now.Day())
	//fmt.Printf("now %dhour %dminute %dsecond.\n", now.Hour(), now.Minute(), now.Second())
	//var army string = "오늘은 $가와 $만에 중서을 다하는 대한민 $육군임니다"
	//armyField := strings.NewReplacer("$", "국")
	//fmt.Println(army)
	//fmt.Println(armyField.Replace(army))
	//in := bufio.NewReader(os.Stdin)
	//fmt.Print("input your name:")
	//name, err := in.ReadString('\n')
	// if err != nil{
	//log.Fatal(err)

	//}else{
	//fmt.Println(name)
	//}

	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)                // 줄바꿈등 제거. 파이썬의 strip 함수와 비슷
	score, _ := strconv.ParseInt(i, 10, 32) // 10진수 정수형(32bit)으로 변환

	var grade string
	if score >= 90 {
		grade = "A"
	} else {
		grade = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다.\n", score, grade)
}
