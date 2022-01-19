package main

import "fmt"

func main() {

	var v [4]int
	fmt.Println(v)
	v[2] = 5
	fmt.Println(v)

	//Slice - composite literal
	// x := type{values}
	x := []int{3, 5, 6, 7, 9}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])
	fmt.Println(x[1])

	fmt.Println("---------------------")

	// range for loop
	for i, v := range x {
		fmt.Println(i, v)
	}

	y := []int{13, 15, 16, 17, 18, 110}
	fmt.Println(y[1])
	fmt.Println(y)
	fmt.Println(y[1:])
	fmt.Println(y[2:4])

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}

	fmt.Println("---------------------")

	//append
	y = append(y, 21, 12, 34, 54, 454, 4)
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

	fmt.Println("---------------------")

	//deleting from slice
	x = append(x[:3], x[5:]...)
	fmt.Println(x)

	//initialise slice using make
	z := make([]int, 4, 8)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	a := [][]int{x, y, z}
	fmt.Println(a)

	fmt.Println("---------------------")

	//map
	m := map[string]int{
		"James":      35,
		"Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnas"])

	// comma OK idiom
	n, ok := m["Barnas"]
	fmt.Println("n :", n)
	fmt.Println("ok :", ok)

	if n, ok := m["Moneypenny"]; ok {
		fmt.Println("This is inside if statement ", n)
	}

	fmt.Println("---------------------")

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, " : ", v)
	}

	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)

	fmt.Println("---------------------")

}
