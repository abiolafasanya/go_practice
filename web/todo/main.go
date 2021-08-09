package main

import (
	"fmt"
	// "html/template"
	"net/http"
)



func main() {
	http.HandleFunc("/", index)

	fmt.Println("Todo App")

	fmt.Println("Golang Serving running")
	http.ListenAndServe(":3000", nil)
}

