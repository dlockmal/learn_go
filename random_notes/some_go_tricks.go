package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// this harms readability, if you have a function like this you might skip typing the variables
	fmt.Println(emptyReturnExample(20, 13))

	// very useful
	closeureExample()

	// this can be useful to check error conditions
	variableDeclarationExample()

	// fallthrough example using switch statements
	fallthroughSwitchExample()

	// this is a way to run an infinite loop with a break condition
	infiniteLoop()
}

// Naked Return
func emptyReturnExample(a int, b int) (addition int, multiplication int, subtraction int) {
	// if you give variable names as return types, you don't need to re-write them in the return statement
	addition = a + b
	multiplication = a * b
	subtraction = a - b
	return
}

// Closure
func closeureExample() {
	// you can use function closures (which is a function that uses the variables defined outside of the function itself)
	// this is quite useful
	x := 2
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

// Variable declaration within the if clause
// I thought this was standard?
func variableDeclarationExample() {
	if i := 2; i < 10 {
		//do sth
		fmt.Println(i)
	}
	// if serr, ok := err.(*json.SyntaxError); ok {
	// do sth
	// }
}

// Fallthrough for switch
func fallthroughSwitchExample() {
	i := 42
	switch {
	// fallthrough needs to be a final statement within the switch block
	// allows you to continue with the next switch statement
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}

	// here the key point is fallthrough also proceeds with the following statement
	// Whether the case is satisfied or not, it executes the next statement
	// please don't do this
	switch 1 {
	case 1:
		fmt.Println("I will print")
		fallthrough
	case 0:
		fmt.Println("I will also print")
	}

}

// Infinite Loop
func infiniteLoop() {
	// we do not need to use while(True) anymore for infinite loops with sleeps or break statements
	for {
		// do sth
		// and break if there's an exit condition
		i := rand.Intn(100)
		fmt.Println(i)
		if i < 3 {
			// fmt.Println(i)
			break
		}

	}
}
