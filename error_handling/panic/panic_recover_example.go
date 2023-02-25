package main

import (
	"fmt"
)

func main() {
	fmt.Println(DivideFunc(20, 10))
	fmt.Println(DivideFunc(10, 0))
}

func DivideFunc(num1, num2 int) int {
	defer func() {
		fmt.Println(recover()) // recover will return nil if there is no panic
	}()
	var q int
	if num2 == 0 {
		panic("Can't divide by zero")
	} else {
		q = num1 / num2
	}
	return q
}
