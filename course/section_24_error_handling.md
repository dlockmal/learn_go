# Section 24: Error handling
## Understanding
- Go does not have exceptions
- the try-catch-finally idiom, results in convoluted code
- multi-value returns make it easy to report an error without overloading the return value
- error handling is quite different but pleasant 
- There is Defer, Panic, and Recover
- The Go Blog has a great article called "Error Handling and Go"
- builtin package "error"
- type error is in interface, an error in go is just an interface

## Checking errors
- ALWAYS check your errors
	- Almost always

Super Basic Example
```
func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
```

Pretty Simple Example
```
func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	r := strings.NewReader("James Bond")
	
	io.Copy(f, r)
}
```

Last Simple Example
```
func main() {
	f, err != os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
```
- almost always check your errors... like above

## Printing and Logging
```
// log and fmt: these are the same
fmt.Println()

// you can also send this somewhere else
log.Println()

// dead, exit code
// deferred functions don't run
log.Fatalln()
	os.Exit()
	
// Not fatal yet, but on that way
// with Panic you can use Recover
// deferred functions run
log.Panicln()

Panic()
```

Log to file in the directory
```
func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	
	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened:", err)
	}
	defer f2.Close()
	
	fmt.Println("check the log.txt file in the directory")
}
```
- probably try to stick to the log package

## Recover

## Errors with info