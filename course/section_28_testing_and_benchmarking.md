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

## Benchmark

## Coverage

## Benchmark examples

## Review