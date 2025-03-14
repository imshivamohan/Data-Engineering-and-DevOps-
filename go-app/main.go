package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on port 8081")
	http.ListenAndServe(":8081", nil)
}

