package main

import (
	"fmt"
	Valuetype "gostruct/struct"
)

func main() {
	var Department = Valuetype.Department{DepartmentName: "IT"}
	var Skills = []string{"Golang", "PHP", "Javascript"}
	var Employee = Valuetype.Employee{Name: "Asdam", Age: 10, Skill: Skills, Department: Department}
	fmt.Println(Employee)
}
