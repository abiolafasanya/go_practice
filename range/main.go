package main

import (
	"fmt"	
)


func main() {

	for i, v := range names{
		fmt.Printf("ID: %v\n => ITEM: %v\n",i ,v)
	}
	fmt.Println("Range Example Running...")

	// sum := 0
	for _, v := range lists{
		// sum +=v
		fmt.Println(v)
	}


	for i, v := range email{
		fmt.Printf("Index: %v => Value: %v\n",i, v)
	}
	fmt.Println("List Example Running...")
}