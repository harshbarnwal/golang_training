package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	marshallPreDefinedDT()
	fmt.Println()
	fmt.Println()

	marshalStructs()

	// unmarshal
	unmarshalData()
	fmt.Println()
}

func marshallPreDefinedDT() {
	color.Green("\nMarshal pre-defined datatypes")
	boolType, _ := json.Marshal(true)
	fmt.Println("bool marshal ->", string(boolType))

	intType, _ := json.Marshal(10)
	fmt.Println("int marshal ->", string(intType))

	floatType, _ := json.Marshal(10.10)
	fmt.Println("float marshal ->", string(floatType))

	stringType, _ := json.Marshal("Harsh")
	fmt.Println("string marshal ->", string(stringType))

	s := []string{"sun", "mon", "tues", "wed", "thurs", "fri", "sat"}
	sliceType, _ := json.Marshal(s)
	fmt.Println("slice marshal ->", string(sliceType))

	m := map[string]int{"jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6, "jul": 7,
		"aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12}
	mapType, _ := json.Marshal(m)
	fmt.Print("map marshal ->", string(mapType))
}

func marshalStructs() {
	color.Green("\nMarshal structs")
	e1 := EmployeeDetails1{Name: "Test Employee", Id: 80, Position: "Digital Marketer"}
	m1, _ := json.Marshal(e1)
	fmt.Println("marshal of struct without json key defined ->", string(m1))

	e2 := EmployeeDetails2{Name: "Harsh Barnwal", Id: 20, Position: "Software Engineer"}
	m2, _ := json.Marshal(e2)
	fmt.Println("marshal of struct with json key defined ->", string(m2))
}

type EmployeeDetails1 struct {
	Name     string
	Id       int
	Position string
}

type EmployeeDetails2 struct {
	Name     string `json:"employee_name"`
	Id       int    `json:"employee_id"`
	Position string `json:"position"`
}

func unmarshalData() {
	color.Magenta("\nUnmarshal data into map")

	byt := []byte(`{"pi":6.13,"place":["New York","New Delhi"]}`)
	var data map[string]interface{}
	if err := json.Unmarshal(byt, &data); err != nil {
		color.Red("Error while unmarshalling data into map", err)
	}
	fmt.Println("unmarshal data ->", data)

	color.Magenta("\nUnmarshal data into struct")
	str := `{"employee_name": "Harsh Barnwal", "employee_id": 20, "position": "Software Engineer"}`
	ed := EmployeeDetails2{}

	err := json.Unmarshal([]byte(str), &ed)
	if err != nil {
		color.Red("Error while unmarshalling data into struct", err)
	}
	fmt.Println(ed)
	fmt.Println("name ->", ed.Name, "id ->", ed.Id, "position ->", ed.Position)
}
