package main

type User struct {
	username string  `json:"username"`
	firstname string  `json:"firstname"`
	lastname string  `json:"lastname"`
	email string  `json:"email"`
	password string  `json:"password"`
}