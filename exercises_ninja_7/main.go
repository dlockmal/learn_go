package main

import "fmt"

func main() {
	exercise_1()
	exercise_2()
}

func exercise_1() {
	// create a value and assign it to a variable
	// Print the address of that value
	x := 42
	fmt.Println("Exercise 1 is stored at:", &x)
}

type person struct {
	name string
}

func exercise_2() {
	// create a person struct
	// create a func called "changeMe" with a *person as a parameter
	// change the value stored at the *person address
	// to dereference a struct, use (*value).field
	// p1.first
	// (*p1).first
	// As an exception, if the type of x is a named pointer type and (*x).f
	// is a valid selector expression denoting a field (but not a method),
	// x.f is shorthand for (*x).f
	// create a value of type person
	// print out the value
	// in func main
	// call "changeMe"
	// in func main
	// print out the value
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	(*p).name = "Miss Moneyp"
}
