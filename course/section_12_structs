# Section 12: Structs

## Struct
- data structure
- aggregate VALUES of different TYPE
- creating a value of type
- composite data structure
- complex data structure
```
type person struct{
	first string
	last string
}

func main() {
	p1 := person{
		first:"James",
		last:"Bond",
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}
	fmt.Println(p1)
	fmt.Println(p2.first, p2.last)
}
```

## Embedded structs
- one type embedded in another type
- you could have a collision in which case you can use
	- sa1.person.first
	- sa1.first
```
package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 32,
		}, 
		lkt: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
}
```

## Reading Documentation
- non-blank field names need to be unique per struct

```
// an empty stuct
struct {}

// A struct with 6 fields
struct {
	x, y int
	u float32
	_ float32 // padding
	A *[]int
	F func()
}
```

## Anonymous Structs
- anonymous because the struct doesn't have a name
- If you want to avoid code pollution
- Only need a struct to use in one little area
```

func main() {
	// remember this is a composite literal
	p1 := struct {
	first string
	last string
	age int
	}{
		first: "James",
		last: "Bond",
		age: 32,
	}
	fmt.Println(p1)
}
```

## Housekeeping
- alias'ing a type is not a good practice
- we create VALUES of a certain TYPE that are stored in VARIABLES
- and those VARIABLES have IDENTIFIERS
```
type person struct {
	first string
	last string
}

// this is valid but don't do it
p1 := person{
	"James"
	"Bon"
}
```
- You don't create classes, you create a TYPE
- You don't instantiate, you create a VALUE of TYPE
- Again ALIAS TYPE isn't the best practice