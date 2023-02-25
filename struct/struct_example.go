package main

import "fmt"

func main() {
	person := Person{fName: "Harsh", sName: "Barnwal", age: 22}
	employee := Employee{personData: person, position: "Software Engineer"}

	person.printDetails()
	employee.printDetails()
}

type Person struct {
	fName string
	sName string
	age   int
}

type Employee struct {
	personData Person
	position   string
}

func (p Person) printDetails() {
	fmt.Println(p)
}

func (e Employee) printDetails() {
	fmt.Println(e)
}
