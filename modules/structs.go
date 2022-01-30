package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func main() {

	//structs
	p1 := person{
		"gaurav",
		"mahal",
		23,
	}
	fmt.Println(p1)

	p2 := person{
		first:  "todd",
		second: "mech",
		age:    33,
	}

	fmt.Println(p2.first, p2.second, p2.age)

	fmt.Println("-----------------")

	//Embedded structs

	type secretagent struct {
		person
		ltk bool
	}

	sa1 := secretagent{
		person: person{
			first:  "James",
			second: "Bond",
			age:    35,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.second, sa1.age, sa1.ltk)

	fmt.Println("-----------------")

	//Anonymous Structs

	a1 := struct {
		first  string
		second string
		age    int
	}{
		first:  "Money",
		second: "Penny",
		age:    30,
	}

	fmt.Println(a1)

	fmt.Println("-----------------")
}
