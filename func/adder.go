package main

func adder() func(int) int{
	var sum = 0
	return func(x int) int{
		sum +=x
		return sum
	}
}