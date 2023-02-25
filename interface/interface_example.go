package main

import (
	"fmt"
)

func main() {
	// with variable type as interface
	var suzukiCar Vehicle
	suzukiCar = Suzuki{model: "S-Cross Alpha AT"}
	suzukiCar.accelerate()
	// fmt.Println(suzukiCar.model) // this will throw error as Vehicle doesn't contains model and Suzuki struct contains this

	var lamboCar Lamborghini
	lamboCar = Lamborghini{model: "AVENTADOR LP 780-4 ULTIMAE"}
	lamboCar.accelerate()
	fmt.Println(lamboCar.model) // this will work as type declaration is Lamborghini
}

type Vehicle interface {
	accelerate()
}

type Suzuki struct {
	model string
}

func (s Suzuki) accelerate() {
	fmt.Println("accelerating at the rate of 40km/hr")
}

type Lamborghini struct {
	model string
}

func (l Lamborghini) accelerate() {
	fmt.Println("accelerating at the rate of 100km/hr")
}
