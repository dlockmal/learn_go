package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println(exercise_1())
	ii := exercise_2(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("The sum of the numbers for Exercise 2 is: ", ii)
	n := exercise_2_1(xi)
	fmt.Println("The sum of the number for Exercise 2.1 is: ", n)
	defer exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
	exercise_7()
	fmt.Println(exercise_8()())
	exercise_9()
	exercise_10()
}

func exercise_1() (string, int) {
	// create a function with the identifier foo that returns an int
	// create a func with the identifier bar that returns an int and a string
	// call and  print both funcs
	x := 1
	y := "Whelp this is a string for Exercise:"
	return y, x
}

func exercise_2(x ...int) int {
	// take in a variadic parameter
	// pass in a value of type []int into your function (unfurl the []int)
	// returns the sum of all values of type int passed in
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func exercise_2_1(x []int) int {
	// takes in a parameter of type []int
	// returns the sum of all values of type int passed in
	// non-variadic parameter
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func exercise_3() {
	// use the "defer" keyword to show that a deferred func runs after the func
	// containing it exists
	fmt.Println("This is a deferred function for Exercise: 3")
}

func exercise_4() {
	// create a user-defined struct with
	// the fields
	// first
	// last
	// age
	// attach a method to type person with
	// the identifier "speak"
	// the method should have the person say their name and age
	// create a value of type person
	// call the method from the value of type person
	x := person{
		"James",
		"Bond",
		50,
	}
	x.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old")
}

func exercise_5() {
	// create a type SQUARE
	// create a type CIRCLE
	// attach a method to each that calculates AREA and returns it
	// create a type SHAPE that defines an interface as anything that has the AREA method
	// create a func INFO which takes the type shape and then prints the area
	// create a value of type square
	// create a value of type circle
	// use func info to print the area of square
	// use func info to print the area of circle
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}
	info(circ)
	info(squa)
}

func info(s shape) {
	fmt.Println("The area of the shape is", s.area())
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func exercise_6() {
	// build and use an anonymous func
	func(x int) {
		fmt.Println("The meaning of life is:", x)
	}(42)
}

func exercise_7() {
	// assign a func to a variable then call that function
	x := func(y int) {
		fmt.Println("The year may or may not be", y)
	}
	x(2018)
}

func exercise_8() func() string {
	// create a function that returns a function
	// assign the returned func to a variable
	// call the returned func
	return func() string {
		return "This is the output from exercise 8"

	}
}

func exercise_9() {
	// pass a func into a func as an argument... a callback
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := odd(sum, ii...)
	fmt.Println("exercise_9 results:", s)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func exercise_10() {
	// create a func which "encloses" the scope of a variable
	{
		y := 42
		fmt.Println("This is exercise_10:", y)
	}
}
