package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1

	fmt.Fprintf(w, "Random Number = %d", num)
}

func main() {
	http.HandleFunc("/random", random)
	fmt.Println("Server running at http://localhost:8080/random")
	http.ListenAndServe(":8080", nil)
}