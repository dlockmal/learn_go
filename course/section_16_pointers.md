# Section 16: Pointers

## What are pointers?
- Points to some location in memory where a value is stored
- Remember the ampersand (&)
- Can't sign an point to an int to an int
- *[type] is a pointer
- *[variable] is dereferencing a memory address
```
func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	// the & operator provides the memory address
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	// above returns  '*int' which is a pointer to an int
	var b *int = &a // or b := &a
	fmt.Println(b)
	// this asterisk is different then a pointer
	// this dereferenes the address
	fmt.Println(*b)
	fmt.println(*&a)
}
```

## When to use pointers?
- when you have a large chunk of data you don't want to pass around your app, you can pass the address
- You need to change something that's at a specific location in memory
- ==By default everything in go is 'pass by value'==
```
func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("xbefore", x)
	foo(x&)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println('y after", y)
	fmt.Println("y after", *y)
}
```

- This may also be  called 'mutate' the value at the address
- All pass by value

## Method sets
- Method sets are the methods attached to a type
- Set of methods attached to a type
- A non-pointer receiver
	- works with values that are POINTERS or NON-POINTERS
- A pointer receiver
	- only works with values that are POINTERS
- more intermediate / advanced topic
- be aware and remember to look at the docs

```
// insert code example here
```