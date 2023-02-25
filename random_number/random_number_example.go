package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("random int", rand.Intn(100))
	fmt.Println("random float", rand.Float64())
	fmt.Println()

	// currently in above technique the number is not totally random, to make it true random we have to provide seeding to the compiler

	rand.NewSource(time.Now().Unix())
	myRand := trueRandom(1, 20)

	fmt.Println("true random int", myRand)

}

func trueRandom(min, max int) int {
	return rand.Intn(max-min) + min
}
