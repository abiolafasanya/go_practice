package main

import (
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type users struct {
	ID int  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname string  `json:"lastname"`
	Email string  `json:"email"`
}

var exec = false

func main() {
	
	// Connect database
	db, err := sql.Open("mysql", "root@tcp(localhost)/golalng")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Database connected successfullly")
	
	// Select from database
	results, err := db.Query("SELECT * FROM users");

	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
	  var user users
	
	  err = results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
	  if err != nil {
		  panic(err.Error())
	  }
	  fmt.Printf("Firstname: %s\n Lastname: %s\n Email: %s\n", user.Firstname, user.Lastname, user.Email)
	

	  fmt.Println("Data Selected Successfullly")
	}

	// insert into database book
		// newUser := new(users)
		// newUser.Firstname = "james"
		// newUser.Lastname = "smith"
		// newUser.Email = "smith@mail.com"
		
		insert, err := db.Query("INSERT INTO users (firstname, lastname, email) VALUES ('abiola', 'fasanya', 'abiola@mail.com')");
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Println("Data Inserted Successfullly")

	

}

