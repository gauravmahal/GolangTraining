### Important links 
* https://gobyexample.com/ 

### Go modules  
```go
go mod init // creates a new module, initializing the go.mod file that describes it.
go build; go test // and other package-building commands add new dependencies to go.mod as needed.
go list -m all // prints the current module’s dependencies.
go get // changes the required version of a dependency (or adds a new dependency).
go mod tidy // removes unused dependencies.
```

### Programming fundamentals

#### String blog
* Strings, bytes, runes and characters in Go - https://go.dev/blog/strings 
* String is in effect a read-only slice of bytes
```go
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //some of the bytes in the sample string are not valid ASCII, not even valid UTF-8
fmt.Println(sample) // ��=� ⌘
fmt.Printf("% x\n", sample) // bd b2 3d bc 20 e2 8c 98
fmt.Printf("%+q\n", sample) // "\xbd\xb2=\xbc \u2318" %q (quoted) verb will escape any non-printable byte sequences in a string
```
* String literals always contain UTF-8 text as long as they have no byte-level escapes in Go and not Go strings. 
* Unicode standard uses the term “code point” to refer to the item represented by a single value
* GO defines rune as an alias for the type int32 
    '⌘' - rune with integer value 0x2318
* for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration.

#### iota 
* It is reset to 0 whenever the reserved word const appears in the source

### Grouping Data

#### Arrays
* Arrays are values. Assigning one array to another copies all the elements.
* In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
* The size of an array is part of its type. The types [10]int and [20]int are distinct.

 * Slices allow you to group together VALUES of same TYPE

 ### Structs

 ### Functions
 * Everything in Go is PASS BY VALUE
 * Variadic parameter has to be the final parameter
 * Empty Interface - Every value can be put in there 
 ```go
 type emptyInetrface interface {}
 ```
 * [Compostion in Go](https://www.ardanlabs.com/blog/2015/09/composition-with-go.html)
 * Nested function declaration is not allowed in Go. But you can assign them to a variable like expression
 ```go
 //Not allowed
func main() {
    func inc(x int) int { return x+1; }
}
//allowed
func main() {
    inc := func(x int) int { return x+1; }
}

 ```

### Poniters
#### Method Sets 
* Method set of any type T consists of all methods declared with receiver type T.
* Method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T


### Application
 * Types, constants, fields with a capital first letters are visible to external program/packages.
 * [About JSON](https://htmlpreview.github.io/?https://github.com/GoesToEleven/golang-web-dev/blob/master/040_json/README.html)
* Go uses these six types for all values parsed into interfaces:
```go
    bool, for JSON booleans
    float64, for JSON numbers
    string, for JSON strings
    []interface{}, for JSON arrays
    map[string]interface{}, for JSON objects
    nil for JSON null
```
* https://mholt.github.io/json-to-go/ 
* crypto/bcrypt - for encrtyption of passwords 