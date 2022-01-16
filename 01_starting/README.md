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

