package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type customErr struct {
	info string
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
}

func exercise_1() {
	// start with some code
	// instead of using a blank identifier, make sure the code is checking and handling the error
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func exercise_2() {
	// start with some code
	// Create a custom error message using "fmt.Errorf"
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Errorf("There has been an error marshaling: %v", err)
	}
	fmt.Println(string(bs))
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(f error) {
	fmt.Println("foo ran -", f, "\n")
}

func exercise_3() {
	// create a struct customErr which implements the builtin.error interface
	// create a func "foo" that has a value of type error as a parameter.
	// create a value of type "customErr" and pass it into "foo"
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

func exercise_4() {
	// start with some code
	// use the sqrt.Error struct as a value of type error
	// use these numbers for your location
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
