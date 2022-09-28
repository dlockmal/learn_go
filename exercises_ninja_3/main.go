package main

import "fmt"

func main() {
	exercise2()
}

// print every number from 1 to 10,000
func exercise1() {
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

// print every code point of the uppercase alphabet three times
func exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

// create a for loop using for condition {} syntax, have it print out the years you have been alive
func exercise3() {
	i := 1984
	for i <= 2022 {
		fmt.Println(i)
		i++
	}
}

// create a loop using for {} and have it print the years you've been alive
func exercise4() {
	i := 1984
	for {
		if i > 2022 {
			break
		}
		fmt.Println(i)
		i++
	}
}

// Print out the remainder which is found for each number between 10 and 100 when it is divided by 4
func exercise5() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

// create a program that shows the "if statement in action"
func exercise6() {
	if 1 == 1 {
		fmt.Println("What did you expect")
	}
}

// Create a program that uses "else if" and "else"
func exercise7() {
	if 1 != 1 {
		fmt.Println("This should never print")
	} else if 1 == 1 {
		fmt.Println("This will print")
	} else {
		fmt.Println("This will not print")
	}
}

// Create a program that uses a switch statement with no switch expression specified
func exercise8() {
	switch {
	case true:
		fmt.Println("Now this will print because the default is true")
	default:
		fmt.Println("This will print but it's pretty boring")
	}
}

// Use a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER "favSport"
func exercise9() {
	favSport := "football"
	switch favSport {
	case "notIt":
		fmt.Println("this isn't going to work")
	case "football":
		fmt.Println("This is it")
	}
}
