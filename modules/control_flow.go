package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i < 3; i++ {
		fmt.Println("Hello everyone")
	}

	//while loop
	i := 0
	for i < 3 {
		fmt.Println("again hello")
		i++
	}

	for v := 32; v < 128; v++ {
		fmt.Printf("%v\t%#U\n", v, v)
	}
	switch i < 5 {
	case true:
		fmt.Println("i < 5")
		fallthrough // made next statement executable
	case false:
		fmt.Println("i >= 5")
	default:
		fmt.Println("default")
	}

}
