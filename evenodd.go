package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func evenodd(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("num"))

	if n%2 == 0 {
		fmt.Fprintf(w, "%d is Even", n)
	} else {
		fmt.Fprintf(w, "%d is Odd", n)
	}
}

func main() {
	http.HandleFunc("/evenodd", evenodd)
	fmt.Println("Server running at http://localhost:8080/evenodd?num=8")
	http.ListenAndServe(":8080", nil)
}