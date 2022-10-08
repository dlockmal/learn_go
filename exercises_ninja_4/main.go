package main

import (
	"fmt"
)

func main() {
	exercise8()
}

func exercise1() {
	//composite literal
	// create array which old 5 values of TYPE INT
	// Assign VALUES to each index position
	// Range over the array and print the values out
	// Use format printing print out the TYPE of the array
	x := [5]int{10, 100, 1000, 10000, 100000}
	for i, v := range x {
		fmt.Println(i, v)
	}
	// %T is the type - good work
	fmt.Printf("%T", x)
	fmt.Println("We should prefer slices over an array")
}

func exercise2() {
	// Using composite literal
	// create a SLICE of TYPE INT
	// Assign 10 values
	// Range over the slice and print the values out
	// Print the TYPE of the slic using format printing
	x := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}

func exercise3() {
	// Using SLICING create the following new slices
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func exercise4() {
	// start with the slice below
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//append to that slice the value 52
	x = append(x, 52)
	// print out slice
	fmt.Println(x)
	// append the values in one line
	x = append(x, 53, 54, 55)
	// print out slice
	fmt.Println(x)
	// append the below slice to x
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	// start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// use APPEND and SLICING to get these values here which you should ASSIGN to a variable "y" then print
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exercise6() {
	// Create a slice to store the names of all of the states in the USA
	// Use MAKE and APPEND to do this
	x := make([]string, 50, 51)
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`,
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`,
		` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`,
		` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	for i, v := range states {
		x[i] = v
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func exercise7() {
	// create a slice of a slice of string
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xp := [][]string{x, y}
	fmt.Println(xp)
	for i, xs := range xp {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\tindex position: %v\tvalue: \t%v\n", j, val)
		}
	}
}

func exercise8() {
	// Create a map with a key of TYPE string which is a person's last_first name, and a value of TYPE []string which stores their favorite things.
	// Store three records  in your map
	// Print out all the values, along with their index position in the slice
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	// exercise 9 is to add a record to the map and print using range again.
	m["McLeod_Todd"] = []string{"Programming Go", "Michigan Hats", "Good lighting setups"}
	fmt.Println("**************************************")

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	// exercise 10 is to delete a record to the map and use range again
	delete(m, "bond_james")
	fmt.Println("**************************************")
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
