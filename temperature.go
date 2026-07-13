package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func temp(w http.ResponseWriter, r *http.Request) {

	c, _ := strconv.ParseFloat(r.URL.Query().Get("c"), 64)

	f := (c * 9 / 5) + 32

	fmt.Fprintf(w, "%.2f Celsius = %.2f Fahrenheit", c, f)
}

func main() {

	http.HandleFunc("/temp", temp)

	fmt.Println("Temperature Converter Service Started...")
	fmt.Println("Open: http://localhost:8080/temp?c=37")

	http.ListenAndServe(":8080", nil)
}