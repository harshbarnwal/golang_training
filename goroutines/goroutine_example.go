package main

import (
	"fmt"
	"sync"
)

func main() {
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

var wg = sync.WaitGroup{}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1 index", i)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2 index", i)
	}
	wg.Done()
}
