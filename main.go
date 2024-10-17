package main

import (
	"fmt"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return fmt.Sprintf("%s%s%d", str1, str2, num)
}
func main() {

	r := ConcatStringsAndInt("r", "e", 5)
	fmt.Println(r)
}
