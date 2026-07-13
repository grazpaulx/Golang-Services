package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func square(w http.ResponseWriter, r *http.Request) {

	num, _ := strconv.Atoi(r.URL.Query().Get("num"))

	fmt.Fprintf(w, "Square = %d", num*num)
}

func main() {

	http.HandleFunc("/square", square)

	fmt.Println("Square Number Service Started...")
	fmt.Println("Open: http://localhost:8080/square?num=5")

	http.ListenAndServe(":8080", nil)
}