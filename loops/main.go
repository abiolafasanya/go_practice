package main

import (
	"fmt"
	"time"
)


func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	for i, j := 0, 1; i < 10 && j < 10; i ,j = i+1, j+1 {
		fmt.Println(i, j)
	}

	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d, %c", i, s[i])
		fmt.Println()
	}

	for i, v := range s{
		fmt.Printf("%d, %c", i,v)
		fmt.Println()
	}
	for _, v := range s{
		fmt.Printf("%c",v)
		fmt.Println()
	}
	
	fmt.Println()
	for i := 20; i <= 100; i = i + 5{
		if i >= 50 && i < 75 {
			fmt.Println("Passed", i)
			}else if i > 74 && i < 100 {
				fmt.Println("Outstanding Passed", i)
		}else if i == 100 {
			fmt.Println("Excellently Passed", i)
		}else{
			fmt.Println("Failed", i)
		}
	}
	
	fmt.Println()

	for _, i := range s {
		if i == ' ' {
			continue
		}
		fmt.Printf("%c", i)
		fmt.Println(string(i))
	}

	fmt.Println("Completed", time.Second)
}