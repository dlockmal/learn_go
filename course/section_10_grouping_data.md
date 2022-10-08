# Section 10: Grouping data

## Array
- zero based index, pretty standard
- arrays aren't really used in Go
	- These are primarily building blocks for slices
	- Use slices instead
- numbered sequence of a single type
```
var x [5]int
fmt.Println(x)
x[3] = 42
fmt.Println(len(x))
```

## Slice - composite literal
- composite data type
```
func main() {
	x := type{values} // COMPOSITE LITERALS
	x := []int{4,5,7,8,42}
}

// a SLICE allows you to group together VALUES of the same TYPE
```

## Slice - for range
```
func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])
	
	// i is the index and v is the value
	// i for index and v for value variable names is common
	for i, v := range x {
		fmt.Println(i, v)
	}
}
```

## Slice - slicing a slice
```
// you can slice a slice with the ':' operator
func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	// prints from position 1 to the end
	fmt.Println(x[1:])
	// will print two digits 1 and 2, up to 3 not including
	// this is also how you delete from a slice
	fmt.Println(x[1:3])
	
	// can iterate over a slice like the following
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
```

## Slice - append to a slice
```
func main() {
	x := []int{1, 2, 3, 5, 8, 13}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	
	y := []int{234, 456, 678, 987}
	// variatic parameter
	// dots afterwards put all the values right there
	x = append(x, y...)
	fmt.Println(x)
}
```

## Slice - deleting from a slice
- https://godoc.org/builtin
	- shows the builtin functions in go 
- it's all about type
- remember the '...' before or after
- why make another function to delete when append is built-in
```
func main() {
	x := []int{4, 5, 7, 8, 42, 77, 88, 99, 987}
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
```

## Slice - make
```
// when you make a slice use the composite literal
// using the make builtin function can save some compute
// the capacity is the length of the underlying array 
// for the slice
func main() {
	// 10 is the length
	// 100 is the capacity
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

```

## Slice - multi-dimensional slice
```
func main() {
 jb := []string{"red", "blue", "green"}
 fmt.Println(jb)
 mp := []string{"cat", "dog", "snake"}
 fmt.Println(mp)
 // multi-dimensional array
 xp := [][]string{jb, mp}
 fmt.Println(xp)
}
```

## Map - introduction
- key value store
- unordered list
- composite literal syntax
```
func main() {
	// m is a good variable for maps
	m := map[string]int{
		// multiple lines should include the trailing comma
		"James":32,
		"Miss Moneypenny":27,
		
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	// if there is not a value for a map the return will be 0
	fmt.Println(m["Barnabas"])
	
	// comma ok idiom
	// check if value exists
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	
	// this is a common pattern
	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the If Print, v")
	}
}
```

- look up via effective go if more information is needed

## Map - add element & range
```
func main() {
	m := map[string]int{
		"james": 32,
		"moneypenny": 27,
	}
	m["todd"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{4,5,7,8,9,42}
	
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
```

## Map - delete

```
func main() {
	m := map[string]int{
		"James": 32,
		"moneypenny": 27,
	}
	delete(m, "James")
	fmt.Println(m)
	
	// it's possible to delete an entry that doesn't exist
	delete(m, "Ian Fleming")
	fmt.Println(m)
	
	// comma ok idiom
	if v, ok := m["moneypenny"]; ok {
		fmt.Println(value", v)
		delete(m, "moneypenny")
	}
	fmt.println(m)
}
```