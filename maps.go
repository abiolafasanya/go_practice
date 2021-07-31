package main

import "fmt"

func main() {
	// emails := make(map[string]string)

	
	emails := map[string]string{
		"Abiola": "abiola@mail.com",
	 	"Sharon": "sharon@mail.com", 
	 	"Esther": "esther@mail.com"}
	
	emails["harbiola"] = "harbiola@gmail.com"
	emails["johndoe"] = "johndoe@gmail.com"
	emails["queen"] = "queen@gmail.com"
	emails["mike"] = "mike@gmail.com"
	

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["harbiola"])
	delete(emails, "mike")
	fmt.Println(emails)
}