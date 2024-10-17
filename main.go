package main

import (
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {

	i, err := strconv.Atoi(a)
	u, err1 := strconv.Atoi(b)

	if err != nil || err1 != nil {
		return 0, fmt.Errorf("invalid input, please provide two integers")
	}
	return i + u, nil
}
func main() {
	fmt.Println(SumTwoIntegers("3.7", "5"))
}
