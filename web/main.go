package main

import (
	"fmt"
	// "database/sql"
    // _ "github.com/go-sql-driver/mysql"
	"net/http"
)

func index (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hi ,this is abiola learning web part of golang</h1> You have requested: %v\n", r.RequestURI)
}

func about (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>This is the about us page</h1> You have requested: %v\n", r.RequestURI)
}

func home(){
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fs := http.FileServer(http.Dir("static/index.html"))

	http.Handle("/home", http.StripPrefix("/static/index.html", fs))


	fmt.Println("Server Starting....")
	http.ListenAndServe(":3000", nil)
}