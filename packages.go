package main

import (
	"fmt"
	"sort"
	"strings"
)


func main() {
	ages := []int{45, 20, 55, 30, 40, 50, 25, 35, 60}
	fmt.Println(sort.SearchInts(ages, 40))
	fmt.Println(ages)

	greeting := "Hello my name is abiola"
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(strings.Index(greeting, "name"))
	fmt.Println(strings.ReplaceAll(greeting, "name", "actual")) 
}