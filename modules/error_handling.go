package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {

	var name, food string

	fmt.Printf("Name ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Food ")
	_, err = fmt.Scan(&food)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, food)
	fmt.Println("--------------------------")

	fmt.Println("Logging")
	f, err := os.Create("error_logging.txt")
	if err != nil {
		log.Fatalln("error opening file:", err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no_file.txt")
	if err != nil {
		log.Println("error opening file:", err)
		// log.Fatalln()
		//  	os.Exit()
		// log.Panicln()
		// 		defer function run
		// 		can use recover

	}
	defer f2.Close()

	fmt.Println("--------------------------")

	fmt.Println("Recover")

	foo()
	fmt.Println("Return normally from f()")

	fmt.Println("--------------------------")

	fmt.Println("Errors with info")

	_, e := sqrt(-10)
	if e != nil {
		fmt.Printf("%+T\n", e)
		log.Println(e)
	}

	fmt.Println("--------------------------")
}

//	Recover
func foo() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	g(0)
	fmt.Println("Returned normally from g()")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Error with info

func sqrt(i float64) (float64, error) {
	if i < 0 {
		return 0, errors.New("Square root of negative number")
	}
	return math.Sqrt(i), nil
}
