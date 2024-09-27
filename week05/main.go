package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	//i := 55
	i := "55"
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))
	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("KIM inha"))
	fmt.Println("i는 %d\n", i)
	fmt.Println("i는", i, "\n")
	fmt.Println("i는", i)
}
