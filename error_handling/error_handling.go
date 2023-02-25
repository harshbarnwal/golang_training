package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	firstElement := 20.0
	sqrtFirst, err := Sqrt(firstElement)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of", firstElement, "is", sqrtFirst)
	}

	secondElement := -10
	sqrtSecond, err2 := Sqrt(float64(secondElement))
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Square root of", secondElement, "is", sqrtSecond)
	}
}

func Sqrt(v float64) (float64, error) {
	if v < 0 {
		var b strings.Builder
		b.WriteString("error while calculating square root of a negative number ")
		b.WriteString(fmt.Sprintf("%v", v))
		return 0, errors.New(b.String())
	}
	return math.Sqrt(v), nil
}
