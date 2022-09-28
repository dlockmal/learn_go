package main

import (
	"fmt"
)

func main() {
	exercise6()

}

func exercise1() {
	x := 42

	// base 10 is decimal
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}

func exercise2() {
	x := 1
	y := 2

	a := x == y
	b := x <= y
	c := x >= y
	d := x != y
	e := x < y
	f := x > y

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

func exercise3() {

	const (
		a        = 42
		b string = "James Bond"
	)

	fmt.Println(a)
	fmt.Println(b)
}

func exercise4() {
	var a int = 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	// this shifts the variable a one bit to the left
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)

}

func exercise5() {
	a := `here is something 
		as a 
		raw string literal
		"you see?"`

	fmt.Println(a)
}

func exercise6() {
	const (
		c = 2018
		d = c + iota
		e = c + iota
		f = c + iota

		g = 2017 + iota
		h = 2017 + iota
		i = 2017 + iota
		j = 2017 + iota
	)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
