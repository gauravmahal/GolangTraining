package main

import "fmt"

func main() {

	var v [4]int
	fmt.Println(v)
	v[2] = 5
	fmt.Println(v)

	//Slices - composite literal
	// x := type{values}
	x := []int{3, 5, 6, 7, 9}
	fmt.Println(x)

}
