package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func factorial(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("num"))

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	fmt.Fprintf(w, "Factorial of %d = %d", n, fact)
}

func main() {
	http.HandleFunc("/factorial", factorial)
	fmt.Println("Server running at http://localhost:8080/factorial?num=5")
	http.ListenAndServe(":8080", nil)
}