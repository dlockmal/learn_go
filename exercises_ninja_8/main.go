package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func exercise_1() {
	// starting with the code below
	// marshal the []user to JSON.
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	// fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("there's an error marshaling the data", err)
		return
	}
	fmt.Println("Exercise 1 data: " + string(bs))

}

func exercise_2() {
	// starting with the code below
	// unmarshal the JSON into a Go data structure
	// json-to-go site
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	bs := []byte(s)
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Exercise 2 data: ", people)
}

func exercise_3() {
	// starting with the code below
	// encode to JSON the []user sending the results to Stdout
	// use json.NewEncoder(os.Stdout).encode(v interface{})
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []person{u1, u2, u3}
	fmt.Println("Exercise 3 data: ")
	// straight to the wire
	json.NewEncoder(os.Stdout).Encode(people)
}

func exercise_4() {
	// starting with the code below, sort the []int and []string for each person
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type ByAge []person
type ByLast []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (bn ByLast) Len() int           { return len(bn) }
func (bn ByLast) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByLast) Less(i, j int) bool { return bn[i].Last < bn[j].Last }

func exercise_5() {
	// starting with the code below sort the []user by
	// age and last
	// also sort each []string "Sayings" for each users
	// print things pleasantly
	u1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []person{u1, u2, u3}

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
	fmt.Println("-------ByAge")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println("-------ByLast")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}

func main() {
	exercise_1()
	exercise_2()
	exercise_3()
	exercise_4()
	exercise_5()
}
