package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Printf("today %dyear %dmonth %dday.\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("now %dhour %dminute %dsecond.\n", now.Hour(), now.Minute(), now.Second())
}
