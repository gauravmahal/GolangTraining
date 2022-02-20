package main

import "fmt"

func main() {
	fmt.Println("Testing Intro")

	fmt.Println("2+3 = ", mySum(2, 3))
	fmt.Println("2+3+5 = ", mySum(2, 3, 5))

	fmt.Println("----------------")

	fmt.Println("----------------")

	fmt.Println("----------------")
}

func mySum(v ...int) int {
	sum := 0
	for _, vi := range v {
		sum += vi
	}
	return sum
}
