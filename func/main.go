package main

import "fmt"

// func greet(n string, f func(string)) {
// 	f(n)
// }

func main() {

	// sum := adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(sum(i))
	// }

	// greet("Abiola", welcome)
	// greet("Abiola", goodBye)


	fname, lname := name("Abiola Fasanya")
	fmt.Printf("Firstname: %v\n Lastname: %v\n",fname, lname)
}