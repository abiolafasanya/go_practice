package main

import (
	"net/http"	
	"text/template"
)


type Todo struct {
	Title     string
	Completed bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var tmpl = template.Must(template.ParseFiles("index.html"))


func index(w http.ResponseWriter, r *http.Request){
	data := TodoPageData{
		PageTitle: "My Todo List",
		Todos: []Todo{
			{Title: "Task 1", Completed: false},
			{Title: "Task 2", Completed: true},
			{Title: "Task 3", Completed: true},
		},
	}
	tmpl.Execute(w, data)
}