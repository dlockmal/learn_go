# Section 14: Functions

## Syntax
- Go is PASS BY VALUE
```
// func is a keyword
// func (r receiver) identifier(parameters) (resturn(s)) {...}

func main() {
	// anything here would be an argument
	foo()
	bar("Darth")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

// we would define parameters here
func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string{
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmtSprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
```

## Variadic Parameter
- don't be mystified by ```...``` used for variadic parameter
	- can look up the lexical elements
- 
```
// func (r receiver) identifier(parameter(s) (return(s)) { code }

func main() {
	x := foo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x)
}

// unlimited number of int's
// the type here is a slice of int
func foo(x ...int) int{
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	// this is more readable and clear
	sum := 0
	
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding",v,"to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
```

## Unfurling a Slice
- if you have multiple parameters and you want a variadic parameter it has to be the final parameter
```
func main() {
	// composite literal
	// this is how you could use a slice of int when you function 
	// is looking for multiple int values
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println(x)
	
	// this will run
	// variadic means 0 or more values
	y := sum()
	fmt.Println(y)
}

func sum(x ...int) int {
	sum := 0
	for i, v := raange x {
		sum += v
	}
	return sum
}
```

## Defer
- ```defer``` is a keyword which will defer the execution of a function until wherever it's being called ends, returns, or panicks

```
func main() {
	// this will get deferred until func main exits
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("bar")
}
```

## Methods
```
// user defined struct
type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk book
}
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
```

## Interfaces & polymorphism
- interfaces allow us to define behavior and do polymorphism
- a value can be of more than one type
	- like the writer interface
	- io.writer
- the below is confusing
- There are also assertions and conversions in here
- Interfaces allows us to interface in another way
	- if you have these methods then you're also "my type"
	- human and secret agent share type speak, which is a human interface
- When you see an empty interface, every value can be put in there
	- refer to the documentation
	- godocs.org
- every value has no methods or an empty interface
- an int is a type of emptry interface
- Bill Kennedy composition (composition with go) with an example
```
// keyword - identifier - type
// memorize this structure ^- 
type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}
// a value can be of more than one type
type human interface {
	speak()
}
func (p person) speak() {
	fmt.Println("I am", s.first, s.last, " - the person speak")
}

func bar(h human) {
	// special case where you switch on type
	switch h.(type) {
		case person:
			fmt.Println("I was passed into bar", h(person))
		case secretAgent
			fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was Passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	p1 := person{
		first: "Dr.",
		last: "Yes"k
	}
	
	// this is polymorphism
	// the bar function can take in many types, person and secret agent
	bar(sa1)
	bar(p1)
	sa1.speak()
	
	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	}
```

## Anonymous Func
- also known as anonymouse self executing function
```
func main() {
	foo()
	
	//anonymous function
	func() {
		fmt.Println("Anonymous func ran")
	}()
	
	// anonymous function with argument
	func (x int) {
		fmt.Println("The meaning of life:", x)
	}(42)
}


func foo() {
	fmt.Println(" foo ran")
}
```

## Func Expressions
- and expression is something that we have that comes down to a value
- functions are first class citizen
	- can do what all the other types can do
	- functions are types like any other type
- assign a function to a variable

```
func main() {
	fmt.Println("Hello, Playground")
	
	func(){
		f := fmt.Println("my first func expression")
	}
	f()
	
	x := func(y int) {
		fmt.println("the yar big brother started watching:", y)
	}
	x(1984)
}
```

## Returning a Func
- return a function from a function
```
// returning a string
func main() {
	s1 := foo()
	fmt.Println(s1)
	
	x := bar()
	fmt.Printf("%T\n", x)
	
	i := x()
	fmt.Println(i)
	// to clean up you could have used
	// fmt.Println(x())
	
	fmt.Println(bar()())

}

func foo() string {
	// this would work just as well
	// return "Hello World"
	s:= "hello world"
	return s
}

func bar() func() int {
	return func() int{
		return 1984
	}
}
```

## Callback
- passing a func as an argument
- function programming not something that is recommended in go, however, it is to be aware of callbacks
- idiomatic go: write clear, simple, readable code
- known as functional programming
```
fucn main() {
	ii := int{1,2,3,4,5,6,7,8}
	s := sum(i...)
	fmt.Println("all numbers", s)
	
	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)
}

func sum(xi ...int) int{
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
```

## Closure
- one scope enclosing other scopes
	- variables declared in the outer scope are accessible in inner scopes
- closure helps us limit the scope of variables
```
func main() {

	{
	y :=42
	fmt.Println(y)
	}
	// you can't print y out of the code block

}
```

```
func main() {
	// these have different scopes and will return different numbers
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x ++
		return x
	}	
}
```

## Recursion
- a func that calls itself
- could do with loops
	- maybe loops over recursion

```
func main() {
	n := factorial(4)
}

// this is recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// this is how you would do it with a loop
// this is much more effecient and easier to read
func loopFact(n int) int{
	total := 1
	for ; n > 0; n-- {
		total += n
	}
	return total
}
```