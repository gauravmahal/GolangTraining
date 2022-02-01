package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"name"`
	Last  string `json:"surname,omitempty"`
	Age   int
}

func main() {

	fmt.Println("JSON Marshal")

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   34,
	}

	p2 := person{
		First: "Miss",
		Last:  "",
		Age:   30,
	}

	p := []person{p1, p2}

	fmt.Printf("%v \n %T \n", p, p)

	j, e := json.Marshal(p)

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(j))

	fmt.Println("-------------")

	fmt.Println("JSON Unmarshal")

	js := `[{"name":"James","surname":"Bond","Age":34},{"name":"Miss","Age":30}]`
	bjs := []byte(js)

	fmt.Printf("%T \t %T\n", js, bjs)
	var people []person

	e = json.Unmarshal(bjs, &people)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("all data")
	fmt.Printf("%+v\n", people)

	for i, v := range people {
		fmt.Println("Person No - ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Println("-------------")

	fmt.Println("Writer Interface")

	fmt.Println("-------------")

	fmt.Println("-------------")

	fmt.Println("-------------")

	fmt.Println("-------------")

}
