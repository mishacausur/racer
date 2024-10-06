package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	array = FiveSteps(array)
	fmt.Printf("%v", array)
}

func FiveSteps(array [5]int) [5]int {
	var reversed [5]int

	for i := 0; i < len(array); i++ {
		reversed[i] = array[len(array)-1-i]
	}
	return reversed
}
