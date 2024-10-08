package main

import "fmt"

func main() {
	input := [9]string{
		"проснуться",
		"позавтракать",
		"сходить в школу",
		"пообедать",
		"погулять с друзьями",
		"сделать домашнюю работу",
		"попрограммировать на Go",
		"поужинать",
		"лечь спать",
	}
	PrettyArrayOutput(input)
}

func ThirdElementInArray(array [6]int) int {
	return array[2]
}

func FindMinMaxInArray(array [10]int) (int, int) {
	min, max := array[0], array[0]

	for i := 1; i < 10; i++ {
		if array[i] < min {
			min = array[i]
		}
		if array[i] > max {
			max = array[i]
		}
	}

	return min, max
}

func SumOfArray(array [6]int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 7; i++ {
		fmt.Printf("%d я уже сделал: %s\n", i+1, array[i])
	}
	for i := 7; i < 9; i++ {
		fmt.Printf("%d не успел сделать: %s\n", i+1, array[i])
	}
}
