package main

import (
	"fmt"
)


func main() {
	name := "abiola"
	if name == "abiola" {
		fmt.Println("Hello Abiola golang")
	}else {
		fmt.Println("Hello visitor golang")
	}

	color := "red"
	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is Blue")
	// case "green":
	// 	fmt.Println("Color is Green")
	default:
		fmt.Println("Color is neither blue|red|green")
	}

	if color == "red" {
		fmt.Println("color is red")
	}else if color == "blue" {
		fmt.Println("color is blue")
	}else {
		fmt.Println("color is neither red|blue")
	}
}