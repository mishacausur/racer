package main

// import (
// 	"fmt"
// )

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	return s.passingScore < (float64(s.solvedProblems) * s.scoreForOneTask)
}
