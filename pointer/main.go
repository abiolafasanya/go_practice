package main

import "fmt"

func main() {
	a := 55
	b := &a
	c := *b
	*b = 35
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("pointer")
}