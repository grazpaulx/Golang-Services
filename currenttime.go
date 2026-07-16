package main

import (
	"fmt"
	"net/http"
	"time"
)

func currenttime(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	fmt.Fprintf(w, "Current Time = %s", now.Format("03:04:05 PM"))
}

func main() {
	http.HandleFunc("/time", currenttime)
	fmt.Println("Server running at http://localhost:8080/time")
	http.ListenAndServe(":8080", nil)
}