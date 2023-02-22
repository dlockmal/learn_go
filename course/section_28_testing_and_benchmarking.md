# Section 28: Testing & Benchmarking

## Introduction
- be in a file that ends with "_test.go"
- put the file in the same package as the one being tested
- be in a func with an signature "func TestXxx("testing.T")"
- Run a test
	- go test
- deal with test failure
	- use the Error, Fail or related methods to signal failure
```
func main() {
	fmt.Println("2 + 3 =", mySum(2,3))
}

func mySum(xi ...int)int {
	sum := 0
	for _, v := range xi {
	sum += v
	}
	// if you add a +1 here your test will fail
	return sum
}
```

create main_test.go

```
package main

func TestMySum(t *testing.T) {
	x := mySum
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
```

## Table tests
- table of data
- series of tests
```
func TestMySum(t *testing.T) {
	type test struct {
		data []int
		answer int
	}
	
	tests := []test{
		test{[]int{21, 21},42},
		test{[]int{1, 1},2},
		test{[]int{3, 4, 5},12},
		test{[]int{-1, 0, 1},0},
	}
	
	for _, v := range tests{
		x := mysum(v,data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
```

- a good idiom is 
	- expected
	- got

## Example tests
- comment above the code is your documentation
- we're going to bring in writing documentation and simuletaneous test your code
	- Using an Example
- The "strings" package has examples which show tests
- go test ./...
- Great Blog Article
	- The Go Blog
	- Testable examples in Go

acdc.go
```
Package acdc asks if you are ready to rock
package acdc

// Sum adds an unlimlited number of values of type int
func sum(n ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
```

main.go
```
func main() {
	fmt.Println(acdc.Sum(2,3))
	fmt.Println(acdc.Sum2,3,4,5,6,7,8,9)
}
```

main_tests.go
```
package acdc

func ExampleSum() {
	fmt.Println(Sum(2,3))
	// Output:
	// 5
}
```

## Golint
- govet reports suspicious constructs
- golint is a linter for source code
- gofmt reformats go source where where golint prints out style mistakes
```
// reports poor formatting style
golint 

// recursively scan
golint ./...

// how to install
go get -u github.com/golang/lint/golint
```

## Benchmark
- Measure the performance of your code
- There are examples in the course content
```
func main() {
	fmt.Println(saying.Greet("James"))
}
```
- need to add benchmark to the testing file

main_test.go
```
package saying

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear James" {
		t.Error("got", s, "expected", "Welcome my dear James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; t < b.N; i++ {
		Greet("James")
	}
}
```
- This will run all benchmarks in the folder
```
go test -bench .

// this is popular to see
-bench=.
```

## Coverage
- Always cover as much as possible
- 100% is near unachievable 
- You can run coverage analysis using go
```
// check how much coverage your tests have
go test -cover

// output file shows the same 
go test -coverprofile c.out 

// show in browser
go tool cover -html=c.out

// learn more
go tool cover -h
```
