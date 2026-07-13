package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func weather(w http.ResponseWriter, r *http.Request) {

	temp, _ := strconv.Atoi(r.URL.Query().Get("temp"))

	if temp >= 35 {
		fmt.Fprintf(w, "Weather Prediction: Very Hot")
	} else if temp >= 25 {
		fmt.Fprintf(w, "Weather Prediction: Sunny")
	} else if temp >= 15 {
		fmt.Fprintf(w, "Weather Prediction: Pleasant")
	} else {
		fmt.Fprintf(w, "Weather Prediction: Cold")
	}
}

func main() {

	http.HandleFunc("/weather", weather)

	fmt.Println("Weather Prediction Service Started...")
	fmt.Println("Open: http://localhost:8080/weather?temp=30")

	http.ListenAndServe(":8080", nil)
}