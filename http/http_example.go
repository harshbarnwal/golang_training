package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/harsh", handleHarsh)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world from go!!")
	fmt.Println("Loaded root")
}

func handleHarsh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Harsh")
	fmt.Println("Loaded /harsh")
}
