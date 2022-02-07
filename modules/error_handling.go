package main

import (
	"fmt"
	"log"
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

	fmt.Println("Recove")

	fmt.Println("--------------------------")
}
