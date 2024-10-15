package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}

func main() {
	man := Person{name: "Anton", age: 31, address: "Krasnogorsk"}
	man.Print()
}
