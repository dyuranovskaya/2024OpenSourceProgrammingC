package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 202444901
	student1.name = "Aidana"
	student1.gpa = 4.0
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 2024556523
	student2.name = "Kimchol"
	student2.gpa = 4.45
	fmt.Println(student2.gpa)
}
