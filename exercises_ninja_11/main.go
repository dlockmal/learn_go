package main

import (
	"encoding/json"
	"errors"
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
		log.Fatal("JSON did not marshal - here's the error:", err)
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

	bs, err := toJSON(p1)
	if err != nil {
		// log.Println(err)
		log.Fatal(err)
		// return
	}

	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	e := errors.New("this is an error")
	fmt.Println(e)
	fmt.Printf("%T\n", e)

	bs, err := json.Marshal(a)
	if err != nil {
		// this works
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
		// return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON: %v", err))

	}
	return bs, nil
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(f error) {
	// redundant new line
	fmt.Println("foo ran -", f, "\n")
	// assertion is different than conversion
	// fmt.Println("foo ran -", e, "\n", e.(customErr).info)
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

// spend more time on composite literal

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
