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

## Example tests

## Golint

## Benchmark

## Coverage

## Benchmark examples

## Review