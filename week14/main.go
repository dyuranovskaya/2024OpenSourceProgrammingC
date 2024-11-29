package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student
	student1.id = 202444901
	student1.name = "Aidana"
	student1.gpa = 4.0
	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 2024556523
	student2.name = "Kimchol"
	student2.gpa = 4.45
	fmt.Println(student2.gpa)
}
