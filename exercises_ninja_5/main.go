package main

import "fmt"

func main() {
	exercise4()
}

func exercise1() {
	// create the struct type person with the following structure
	// create two values of type person
	// print out the values, ranging over the elements in the slice
	type person struct {
		fname    string
		lname    string
		iceCream []string
	}

	p1 := person{
		fname:    "Darth",
		lname:    "Vader",
		iceCream: []string{"rocky road", "mint choc chip"},
	}
	p2 := person{
		fname:    "Luke",
		lname:    "Skkywalker",
		iceCream: []string{"vanilla", "green tea", "oreo"},
	}

	fmt.Println(p1.fname)
	fmt.Println(p1.lname)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.fname)
	fmt.Println(p2.lname)
	for i, v := range p2.iceCream {
		fmt.Println(i, v)
	}
	//exercise 2
	// store the values of type person in a map with the key of last name
	// access each value in the map
	// print out the values ranging over the slice
	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}
	for _, v := range m {
		fmt.Println(v.fname)
		fmt.Println(v.lname)
		for i, val := range v.iceCream {
			fmt.Println(i, val)
		}
		fmt.Println("-------------------")
	}
}

func exercise3() {
	//create a new type vehicle of struct
	// fields should be doors and color
	type vehicle struct {
		doors int
		color string
	}
	// create two types, truck and sedan with additition field and include the vehicle struct
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}
	// create an item of each type sedan and truck
	v1 := truck{vehicle: vehicle{doors: 4,
		color: "red"},
		fourWheel: true,
	}
	v2 := sedan{vehicle: vehicle{doors: 4,
		color: "black"},
		luxury: true,
	}
	// print the values specifically from each item
	fmt.Println(v1.doors, v1.vehicle.color, v1.fourWheel)
	fmt.Println(v2)
}

func exercise4() {
	// create and use an anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
}
