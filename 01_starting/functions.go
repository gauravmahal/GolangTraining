package main

import "fmt"

func main() {

	r1, r2 := foo("James")
	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println("-------------------")

	//variadic parameters
	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(s)

	fmt.Println("-------------------")

	//Defer

	fmt.Println("-------------------")
}

//func (r reciever) identifier(parameters) (return(s)) { ... }

func foo(s string) (string, bool) {
	fmt.Println("Hello", s, "from foo func")
	return fmt.Sprint(s, " returned from func"), false
}

func sum(v ...int) int {
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	sum := 0
	for _, i := range v {
		sum += i
	}
	return sum
}
