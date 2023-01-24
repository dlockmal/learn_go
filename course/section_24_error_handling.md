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
- The Go Blog: Defer, Panic, and Recover
### Defer
- List itemDefer can be used to do things like ensure the file is close right after opening in
- a deferred function's arguments are evaluated when the defer statement is evaluated
- Deferred function calls are executed in last in first out order after the surrounding function
- Deffered functions may read and assign to the returning function's named return value (name return value: gross)
	
### Panic / Recover
- Recover is only useful inside deferred functions
- a call to recover will only return a value during a panic
- the json package has a real-world example of panic / recover

Example of Panic / Recover (intermediate to advanced)
```
func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
}
```

## Errors with info
```
func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, err) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}
```

There is also this format
```
var ErrNorgateMath = errors.New("norgate math: square root of a negative number")

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}
```

Another Format (error f format printing)
```
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math again: %v", f)
	}
	return 42, nil
}
```

- if we have a type, like a struct, and has a method with Error() string then any value of this type will implicitly implement the error interface (Builtin)

```
type norgateMathError struct {
	lat string
	long string
	err error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}

```

- errors are just another type, you can create values of that type