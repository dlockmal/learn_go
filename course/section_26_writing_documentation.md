# Section 26: Writing Documentation

## Introduction
We need understand how to read documentation first
- godoc.org (this doesn't exist anymore)
- Now it's https://go.dev
	- I like /doc/effective_go
	- https://go.dev/doc/effective_go
- golang.org and godoc.org are now the same
- there is a tool called 'godoc'
	- extracts and generates documentation for go packages
- you can also launch a local webserver that will bring in golang.org
- there is also the command 'go doc'
	- show documentation for package or symbol

## go doc
- allows us to read documentation at the terminal
```
go help doc
go  doc
go doc -cmd cmd/doc
go doc fmt
go doc fmt.Printf
go doc json
go doc json.Number
go doc json.Number.Int64
cd go/src/encoding/json; go doc decode
```

## godoc
- you can run the documentation site locally 
- using godoc with spaces works too
	- pick one you like and stick to it
```
godoc -http=:8080
or 
godoc -http :8080

godoc fmt Printf

godoc -src fmt Printf

godoc fmt
```

## godoc.org
File Structure is something similar to

- 02
	- mymath
		- main.go
		- main_test.go
	- main.go (This imports the mymath function)

godoc.org
if you paste your github url in the godoc.org package it will add it... this is wild
- there is also a refresh button on the godoc.org site

## Writing Documentation
- can be generated
- tightly coupled with code
- provide comments above your code
- godoc - package "errors" is a good example
- Write in complete sentances with periods
- Stuff not exported don't get added to the documentation
- Also look at the "fmt" package
	- This has a whole lot of information
	- create doc.go for package documentation
	- all comments in MD turns into the top section of the package documentation

```
// Package mymath provides ACME inc math solutions
package mymath

// Sum adds an unlimlited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
```

- Go Blog
	- Godoc: documenting Go code
	- godoc [full path] (URL or path in your src directory)
- Create documentation for your code, create good comments
- https://blog.golang.org/godoc-documenting-go-code
