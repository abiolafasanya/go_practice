package main

import "fmt"

func main() {
	
	facebook := Facebook{
		username: "harbiola",
		firstname: "Abiola",
		lastname: "Fasanya",
		email: "harbiola78@gmail.com",
	}

	google := Google{
		username: "harbiola",
		firstname: "Abiola",
		lastname: "Fasanya",
		email: "harbiola78@gmail.com",
	}

	fmt.Println(verify(facebook))
	fmt.Println(verify(google))
	fmt.Printf("Facebook Details\n Facebook Username: %v\n Firstname: %v\n Lastname: %v\n Email: %v\n", 
	facebook.username, facebook.firstname, facebook.lastname, facebook.email)
}