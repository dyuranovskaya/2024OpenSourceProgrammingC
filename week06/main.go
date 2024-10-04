package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 55
	f := 12.9
	fmt.Println("value i :%d,value f:%f\n", i, f)
	//fmt.Println("%d*%f=%f/n", i, f, i*f)
	fmt.Println("%d*%f=%f/n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(f))

}
