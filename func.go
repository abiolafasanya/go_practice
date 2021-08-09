package main

import (
	"fmt"	
	"math"
)


func sayHello(n string){
	fmt.Printf("Hello %v\n", n)
}

func sayBye(n string){
	fmt.Printf("GoodBye %v\n", n)
}

func customGreet(n []string, f func (string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64{
 		return (math.Pi * r * r)
}

func main() {
	fmt.Println("function practice, hello abiola keep building")

	// sayHello("Abiola")
	// sayHello("Fasanya")
	// sayBye("Abiola")

	// v := []string{"james", "scott", "simran"}
	// customGreet(v, sayHello)
	// customGreet(v, sayBye)

	// Area of Cirlce
	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Printf("circle 1 is\n%0.3f\nand circle 2 is\n%0.3f\n", a1, a2)
}