### Important links 
* https://gobyexample.com/ 

### Go modules  
```go
go mod init // creates a new module, initializing the go.mod file that describes it.
go build; go test // and other package-building commands add new dependencies to go.mod as needed.
go list -m all // prints the current module’s dependencies.
go get // changes the required version of a dependency (or adds a new dependency).
go mod tidy // removes unused dependencies.
go run  // compile and run the named main go package 
        // --race to find any race condition in program 
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

### Concurrency
* _**Do not communicate by sharing memory; instead, share memory by communicating**_
* runtime package - for curr architecure, CPU etc.
* sync package - provides basic synchronization primitives like Mutex, WaitGroup
* atomic package - provides low-level atomic memory primitives useful for implementing synchronization algorithms

### Channels
* **Channels block**
    * they are like runners in a relay race
* In general, don't use buffers instead make use of hand shaking by passing sender and reciever
* Channel type. Read from left to right.
    * recieveing from channel an int (<-ch int)
    * channel sending an int (ch<- int)
* **Range clause on channel block till channel is close**
* time package -  for time and sleep operations 
* Rob Pike code source [link](https://talks.golang.org/2012/concurrency.slide#25)
* Nice blog:- [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* Context package - used for easy to pass request-scoped values in goroutines
    * [Go Concurrency Patterns: Context](https://go.dev/blog/context)

### Error Handling
* 


#### TODO
    // read go channel documentation
    // select documentation
    // context documentation