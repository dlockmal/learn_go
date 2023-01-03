# Section 18: Application

## JSON Documentation
- golang.org/docs/ > package documentation
	- these are all the packages that are available withing go
	- also godoc.org/encoding/json
- you check your errors right after the place where the activity takes place
	- error checking takes place on a regular cadence
- slice of bytes is commonly used with strings
- keep an eye out in the documentations for * and & to determine if you need to pass pointers to the new decoder func, or similar

## JSON Marshal
- produces JSON
- The fields in the struct are required to have the first letter capitalized
```
type person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := person{
		First: "James",
		Last: "Bond",
		Age: "32", 
	}
	
	p2 := person{
		First: "Miss",
		Last: "Moneypenny",
		Age: 27,
	}
	
	people := []person{p1, p2}
	fmt.Println(people)
	
	bs, err := json.Marshal(people)
	if err = nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	
}
```


## JSON Unmarshal
- github.com/goestoeleven - Todd McLeod
	- golang-web-dev
	- rawgit.com
	- website: ==json to go==
		- gives you a slice of struct
```
// tags are used below, this is cool
type person []struct {
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	
	// people := []person{}
	
	// this is probably more readable but does the same as above
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", people)
	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

```
- marshal vs unmarshal
- encode and decode
	- doing json operation straight from the wire
	- don't have to assign to a variable
	- you can encode it straight to a client/web connection or file
	- uses readers and writers

## Writer interface
- godoc.org/io
- anywhere you have type: writer method attached you can use it where it's asked for, because it's an interface
- the file function has type writer method, if you wanted to you could pass the file to newencoder 
- Reader has it's own interface also
- Println writes to standard output
- Fprintln - used to print to a file
- This is confusing AF
```
func main() {
// println calls Fprintln
 fmt.Println("Hello World")
 fmt.Fprintln(os.Stdout, "Hello World")
 io.WriteString(os.Stdout, "Hello, World")
}

```

## Sort
- golang.org/doc/ > package documentation > search sort
- a slice is a pointer, a length, and a capacity
```
func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "MoneyPenny", "Dr. No"}
	
	fmt.Println(xi)
	fmt.Println(xs)
	
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
	
	sort.
}
```


## Sort custom
- there is a sort package
```
type person struct {
	first string
	age int
}

type ByAge []person
type ByName []person

func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) {a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (bn ByName) Len() int { return len(bn) }
func (bn ByName) Swap(i, j int) {bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].first < bn[j].first }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}
	
	people := []person{p1, p2, p3, p4}
	
	
	sort.Sort(ByAge(people))
	sort.Sort(ByName(people))
	fmt.Println(people)
	
}
```

## bcrypt
```
go get golang.org/x/crypto/bcrypt
```


```
func main() {
	s := 'notarealpassword'
	bs, err := bcrypt.GenerateFromPassword(s[]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	
	loginPword1 := 'notarealpassword'
	
	err = bcrypt.CompareHashandPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're Logged In")
}

```