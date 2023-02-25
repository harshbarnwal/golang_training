package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// adding employees
	emp = append(emp, Employee{Id: "1", FirstName: "Harsh", SecondName: "Barnwal", Address: &Address{City: "Lucknow", State: "Uttar Pradesh"}})
	emp = append(emp, Employee{Id: "2", FirstName: "Test", Address: &Address{City: "Lucknow", State: "Uttar Pradesh"}})

	// fetch all employees
	router.HandleFunc("/employee", GetEmployeeEndpoint).Methods("GET")

	// fetch an employee using employee id
	router.HandleFunc("/employee/{id}", GetEmployeeById).Methods("GET")

	// add an employee
	/*
		here we can post
		{"first_name": "Post", "second_name": "man", "address": {"city": "api", "state": "post"}, "extra": "data"}
	*/
	router.HandleFunc("/employee/{id}", AddEmployee).Methods("POST")

	// delete an employee
	router.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

type Employee struct {
	Id         string   `json:"id",omitempty`
	FirstName  string   `json:"first_name",omitempty`
	SecondName string   `json:"second_name"`
	Address    *Address `json:"address",omitempty`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state",omitempty`
}

var emp []Employee

func GetEmployeeEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(emp)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range emp {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Employee
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = params["id"]
	emp = append(emp, person)
	json.NewEncoder(w).Encode(emp)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var found bool = false
	for index, item := range emp {
		if item.Id == params["id"] {
			found = true
			emp = append(emp[:index], emp[index+1:]...)
			break
		}
	}
	if !found {
		json.NewEncoder(w).Encode("Employee not found")
	} else {
		json.NewEncoder(w).Encode(emp)
	}
}
