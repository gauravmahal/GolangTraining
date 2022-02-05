package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	First string `json:"name"`
	Last  string `json:"surname,omitempty"`
	Age   int
}

// Sort custom ------------------
type shortPerson struct {
	first string
	age   int
}

type byName []shortPerson

func (p byName) Len() int {
	return len(p)
}

func (p byName) Less(i, j int) bool {
	return p[i].first < p[j].first
}

func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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

	fmt.Printf("%+v \t%T \n", p, p)

	j, e := json.Marshal(p)

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(j))

	fmt.Println("-------------")

	fmt.Println("JSON Unmarshal")

	js := `[{"name":"James","surname":"Bond","Age":34},{"name":"Miss","Age":30}]`
	bjs := []byte(js)

	fmt.Printf("js type: %T;\t bjs type: %T\n", js, bjs)
	var people []person

	e = json.Unmarshal(bjs, &people)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Printf("all data: %+v\n", people)

	for i, v := range people {
		fmt.Println("Person No -", i, ": ", v.First, v.Last, v.Age)
	}

	fmt.Println("-------------")

	fmt.Println("Writer Interface")

	// An Encoder writes JSON values to an output stream.
	// File implements both io.Reader and io.Writer interface
	name := "application_read_write.txt"
	{
		// Create file
		f, e := os.Create(name)
		check(e)
		defer f.Close()

		// Get encoder
		en := json.NewEncoder(f)

		en.Encode(p)
		fmt.Println("Encoded json to the file stream opened at", name)

		f.Sync()
		_, e = f.Seek(0, 0)
		check(e)

		by := make([]byte, 100)
		le, e := f.Read(by)
		checkEOF(e)

		fmt.Println("Read from file name:", name, ":", string(by[:le]))
	}
	{

		f, e := os.Open(name)
		check(e)
		de := json.NewDecoder(f)

		t, e := de.Token()
		check(e)

		fmt.Printf("%T: %v\n", t, t)
		for de.More() {
			var val person
			e = de.Decode(&val)
			checkEOF(e)
			fmt.Printf("%+v\n", val)
		}
		t, e = de.Token()
		check(e)

		fmt.Printf("%T: %v\n", t, t)
	}
	fmt.Println("-------------")

	fmt.Println("Sort custom")
	{
		p1 := shortPerson{"Q", 64}
		p2 := shortPerson{"Moneypenny", 27}
		p3 := shortPerson{"M", 56}
		p4 := shortPerson{"James", 32}

		p := []shortPerson{p1, p2, p3, p4}

		fmt.Printf("%v\n", p)
		sort.Sort(byName(p))
		fmt.Printf("%v\n", p)
	}
	fmt.Println("-------------")

	fmt.Println("bcrypt")
	{
		paswrd := `hello`

		bs, e := bcrypt.GenerateFromPassword([]byte(paswrd), bcrypt.MinCost)
		check(e)
		fmt.Println(bs)

		// // one way of getting input
		// val := make([]byte, 20)
		// os.Stdin.Read(val)
		// fmt.Printf("%T:%v", val, string(val))

		// // second way of getting output
		// buf := bufio.NewReader(os.Stdin)
		// st, e := buf.ReadString('\n')
		// check(e)
		// fmt.Println(st)

		// third way of getting input
		scanner := bufio.NewScanner(os.Stdin)
		_ = scanner.Scan()
		enterPasswrd := scanner.Bytes()
		e = bcrypt.CompareHashAndPassword(bs, enterPasswrd)
		if e != nil {
			fmt.Println("Login Failed")
		} else {
			fmt.Println("Login Sucessful")
		}
	}
	fmt.Println("-------------")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkEOF(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}
