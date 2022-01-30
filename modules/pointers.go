package main

import "fmt"

func main() {
	fmt.Println("Poniters ")

	a := 42
	fmt.Printf("value - %d; Type - %T\n", a, a)
	fmt.Printf("value - %v; Type - %T\n", &a, &a) // & gives you the address

	fmt.Println(*&a)

	fmt.Println("---------------")

	fmt.Println("When to use pointer")

	foo(&a)
	fmt.Println(a)

	fmt.Println("---------------")

	fmt.Println("Method Sets")

	fmt.Println("---------------")
}

func foo(x *int) {
	*x++
}
