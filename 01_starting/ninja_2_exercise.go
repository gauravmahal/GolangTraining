package main

import "fmt"

const (
	curr      = 2022 + iota
	next      = 2022 + iota
	afternext = 2022 + iota
	fin       = 2022 + iota
)

func main() {
	a := 42
	fmt.Printf("%d\t%#b\t%#x\n", a, a, a)
	b := 42
	fmt.Printf("%d\t%#b\t%#x\n", b, b, b)
	c := b << 1
	fmt.Printf("%d\t%#b\t%#x\n", c, c, c)
	str := `This is a string 
	literal. 
	Hurray which can write to "several lines"
	!!!`
	fmt.Println(str)

	fmt.Println(curr)
	fmt.Println(next)
	fmt.Println(afternext)
	fmt.Println(fin)
}
