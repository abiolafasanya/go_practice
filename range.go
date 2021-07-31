package main

import "fmt"

func main() {
	lists := []int{1,2,3,4,5,6,7,8,9}

	// for i, list := range lists {
	// 	fmt.Printf("%d - list: %d\n", i,list)
	// }

	for _, list := range lists {
		fmt.Println(list)
	}

	sum :=0
	for _, id := range lists {
		sum +=id
	}
	// fmt.Println("sum", sum)

	// range with map
	emails := map[string]string{
		"Abiola": "abiola@mail.com",
	 	"Sharon": "sharon@mail.com", 
	 	"Esther": "esther@mail.com"}

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}
}