package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var now time.Time = time.Now()
	//fmt.Printf("today %dyear %dmonth %dday.\n", now.Year(), now.Month(), now.Day())
	//fmt.Printf("now %dhour %dminute %dsecond.\n", now.Hour(), now.Minute(), now.Second())
	//var army string = "오늘은 $가와 $만에 중서을 다하는 대한민 $육군임니다"
	//armyField := strings.NewReplacer("$", "국")
	//fmt.Println(army)
	//fmt.Println(armyField.Replace(army))
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString("\n")
	fmt.Println(i, err)
}
