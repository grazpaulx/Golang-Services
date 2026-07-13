package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculator(w http.ResponseWriter, r *http.Request) {

	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	operation := r.URL.Query().Get("operation")

	if operation == "add" {
		fmt.Fprintf(w, "Addition = %d", a+b)
	} else if operation == "sub" {
		fmt.Fprintf(w, "Subtraction = %d", a-b)
	} else if operation == "mul" {
		fmt.Fprintf(w, "Multiplication = %d", a*b)
	} else if operation == "div" {
		if b == 0 {
			fmt.Fprintf(w, "Cannot divide by zero")
		} else {
			fmt.Fprintf(w, "Division = %d", a/b)
		}
	} else {
		fmt.Fprintf(w, "Invalid Operation")
	}
}

func main() {

	http.HandleFunc("/calculator", calculator)

	fmt.Println("Calculator Service Started...")
	fmt.Println("Open: http://localhost:8080/calculator?a=10&b=5&operation=add")

	http.ListenAndServe(":8080", nil)
}