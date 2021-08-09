package main

import (
	"fmt"	
	"strings"
)


func main() {
	a := "abiola fasanya"
	fmt.Println(a[0])
	fmt.Printf("%c\n", a[0])
	fmt.Printf("%c\n", a[3])
	fmt.Println(a[:7])
	fmt.Println(a[3:])
	fmt.Println(a[3:7])
	fmt.Println()

	a += " again "
	a = a + "again "
	a = a + `new:"again"`
	a = a + `new:'cool'`

	fmt.Println(a)
	fmt.Println()

	fmt.Println(len(a))
	fmt.Println()

	abc := "abc"
	b := []byte(abc)
	fmt.Println(abc, b)
	fmt.Printf("%v, %s\n",abc, b)
	fmt.Println()

	abc2 := string(b)
	fmt.Println(abc2)
	fmt.Println()
	fmt.Println("Hello World")

	fmt.Println()
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(a))
	// fmt.Println(strings.HasPrefix(a))
	// fmt.Println(strings.HasSuffix(a))
}