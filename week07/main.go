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

	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input your score : ")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	score, _ := strconv.ParseInt(i, 16, 32)
	if score >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", score)
	}

}
