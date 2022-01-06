package main

import "fmt"

var xg int = 42
var yg string
var zg bool

type ownType int

var xo ownType
var yo int

func main() {
	x := 42
	fmt.Println(x)
	y := "James Bond"
	fmt.Println(y)
	z := true
	fmt.Println(z)
	fmt.Printf("%d\t%s\t%t\n", x, y, z)

	fmt.Println("Global values")
	fmt.Println(xg, yg, zg)
	yg = "James Bond"
	zg = true
	s := fmt.Sprintln(xg, yg, zg)
	fmt.Print(s)

	fmt.Printf("%d\n", xo)
	fmt.Printf("%T\n", xo)
	xo = 42
	fmt.Printf("%d\n", xo)

	yo = int(xo)
	fmt.Println(yo)
	fmt.Printf("%T\n", yo)

	// #6 Quiz
	// https://docs.google.com/forms/d/e/1FAIpQLSfyN4xMJZPoz_2bVy-BXctXfb1a64n4deYF9jj6JLnhpwA3dw/viewscore?viewscore=AE0zAgBxvhwMuxKQW3e0VcTNAhpR3ZoORkVODo_X2Avf53Bk_8ogOsv_jLoBywSdqDhaNG4

}
