package main

import "fmt"

type users struct {
	name string
	age int
	email string
}

var user1 = []users{
	{name: "samuel", age: 25, email: "samuel@gmail.com"},
	{name: "james", age: 25, email: "james@gmail.com"},
	{name: "mike", age: 25, email: "mike@gmail.com"},
}

func main() {
	user := new(users)
	user.name = "abiola"
	user.age = 25
	user.email = "abiola@mail.com"


	fmt.Println("Hello learning struct")
	fmt.Println("Result", user1)
	fmt.Printf("Result", user)
	fmt.Printf("Name : %s\n Age: %d\n Email: %s\n", user.name, user.age, user.email)
}