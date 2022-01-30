package main

import (
	"fmt"
)

func main() {

	r1, r2 := foo("James")
	fmt.Println(r1)
	fmt.Println(r2)

	fmt.Println("-------------------")

	// variadic parameters
	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(s)

	fmt.Println("-------------------")

	// Defer
	// Used for openeing and closing file's etc
	// execution is deferred to the moment the surrounding function returns
	defer bar()

	fmt.Println("-------------------")

	// Methods
	sg := secretagent{
		person: person{
			"James",
			"Bond",
		},
		skg: true,
	}

	sg.speak()

	// Interfaces and Inheritance
	// In Go, values can be of more than one type

	p := person{
		"Miss",
		"Moneypenny",
	}
	caller(sg)
	caller(p)

	fmt.Println("-------------------")

	// Anonymous func

	func(x int) {
		fmt.Println("Anonymous func with int", x)
	}(23)

	fmt.Println("-------------------")

	// Func Expession

	f := func(x int) {
		fmt.Println(x, "passed to a func expression")
	}
	f(12)

	fmt.Println("-------------------")

	// Returning a function

	fr := funcReturn()
	fmt.Println(fr())

	fmt.Println(funcReturn()()) // interesting

	fmt.Println("-------------------")

	// Callback
	es := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)

	fmt.Println("even sum", es)

	fmt.Println("-------------------")

	// Closure

	c1 := incrementor()
	c2 := incrementor()

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c2())

	fmt.Println("-------------------")
}

//func (r reciever) identifier(parameter(s)) (return(s)) { ... }

func foo(s string) (string, bool) {
	fmt.Println("Hello", s, "from foo func")
	return fmt.Sprint(s, " returned from func"), false
}

// -------------

func sum(v ...int) int {
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	sum := 0
	for _, i := range v {
		sum += i
	}
	return sum
}

// -------------

func bar() {
	fmt.Println("defer bar")
}

// -------------

type person struct {
	first  string
	second string
}

type secretagent struct {
	person
	skg bool
}

func (s secretagent) speak() {
	fmt.Println("I am " + s.first + s.second + " secretagent")
}

// -------------

func (s person) speak() {
	fmt.Println("I am " + s.first + s.second + " person")
}

type human interface {
	speak()
}

func caller(h human) {
	fmt.Println("Inside caller ", h)
	switch h.(type) {
	case secretagent:
		fmt.Println("Inside caller assertion -------- ", h.(secretagent).first)
	case person:
		fmt.Println("Inside caller assertion -------- ", h.(person).first)
	}
}

// -------------

func funcReturn() func() int {
	return func() int {
		return 123
	}
}

// -------------

func evenSum(f func(x ...int) int, v ...int) int {
	var ev []int

	for _, i := range v {
		if i%2 == 0 {
			ev = append(ev, i)
		}
	}
	return f(ev...)
}

// Closure --------------

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
