package main

import (
	"fmt"
)

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Println("Employee:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Total Salary: %.2f", e.bonus+e.salary)
}

func main() {
	man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
	man.CalculateTotalSalary()
}
