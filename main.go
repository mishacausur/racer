package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

type Dog struct{}

func (c Dog) MakeSound() {
	fmt.Println("Гав!")
}
