package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	r1 := []rune(strings.ToLower(str1))
	r2 := []rune(strings.ToLower(str2))

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	return string(r1) == string(r2)
}
func main() {
	fmt.Println(AreAnagrams("Кабан", "банка"))
}
