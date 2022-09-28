# Section 8: Control Flow
## Understanding Control Flow
- How your computer reads the program

## Loop - init, condition, post
- ```bit hacking with go by vladimir``` article
-  https://gobyexample.com
-  There is no "while" in go
```
// example
for init; condition; post {}

package main

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
```

## Loop - nesting Loops

```
package main

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("The inner loop: \t%d\n", j)
		}
	}
}
```

## Loop - for statement
- The go programming language specs
- EBNF (Extended Backus Naur Form)
- Effective Go (is this still around)

```
// this is how you would do a while statement
package main

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
```


```
// this method shows a loop that will run indefinitely until there is a break
package main

func main() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
```

## Loop - break & continue
- continue is used to exit a loop while continuing the next iteration
- break is used to stop the service or process
```
func main() {
	x := 0
	for {
		x++
		if x <= 100 {
			break
		}
		if x % 2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
```

## Loop - printing ascii
This was just an example
```
package main

import "fmt"

func main() {
	for x := 33; x <= 122; x++ {
		if x <= 145 {
			fmt.Printf("%#U\n", x)
			x++
		}
	}

}

```


## Conditional - if statement
 - Nothin new to comment
 
 ## Conditional - if, else if, else

```
func main(){
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was  41")
	} else {
		fmt.Println("our value was not 40 or 41")
	}
}
```
 
 ## Loop, conditional, modulus
 - Hands on exercise
```
func main() {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	} 
}
```
 
 ## Conditional - switch statement
 - ==sequential (default), iterative, conditional==
 - no default fall through
 ```
 func main() {
 	switch {
		case false:
			fmt.Println("this should not print")
	}
	switch {
		case (2==4):
			fmt.Println("this should not print 2")
	}
	switch {
		case (3 ==3):
			fmt.Println("prints")
	}
	switch {
		case (4 == 4):
			fmt.Println("also tru, but does not print")
	}
 }
 ```
 
 - You can specify fall through
```
 func main() {
 	switch {
		case false:
			fmt.Println("this should not print")
	}
	switch {
		case (2==4):
			fmt.Println("this should not print 2")
	}
	switch {
		case (3 ==3):
			fmt.Println("prints")
			fallthrough
	}
	switch {
		case (4 == 4):
			fmt.Println("also tru, but does not print")
	}
 }
```
 
 - this one is a bit different
```
 func main() {
 	switch {
		case false:
			fmt.Println("this should not print")
	}
	switch {
		case (2==4):
			// now this will print
			fmt.Println("this should not print 2")
			fallthrough
	}
	switch {
		case (3 == 3):
			fmt.Println("prints")
	}
	switch {
		case (4 == 4):
			fmt.Println("also tru, but does not print")
	}
 }
```
 - ==just don't use fallthrough==
 - there is a default case where if nothing matches true then it defaults to this case
```
case (1 == 2):
	fmt.Println("this is false"
default:
	fmt.Println("this is default")
```
 
 - can switch on strings
 - can also switch on a variable assigned to a value

```
func main() {
	switch "Bond" {
		case "Moneypenny":
			fmt.Println("miss money")
		case "Bond":
			fmt.Println("bond james")
	}
}
```

- multiple cases can be used with commas
```
func main() {
	n := "Bond"
	switch n {
		case "Moneypenny", "Bond", "Do No":
			fmt.Println("miss money or bond or dr no")
	}
}
```
## Conditional - switch statement documentation
- there are expression switches and type switches
- cases contain expressions that are compared against the value of the switch expression
- https://golang.org/ref/spec
- missing switch expression defaults to true
- switch and case are keywords
- cases are evaluated top to bottom until a match is found
- it's idiomatic to wriate an if-else-if-else chain as a switch
 
 ## Conditional logic operators
 ```
 //true
 true && true
 //false
 true && false
 //true
 true || true
 //true
 trute || false
 //false
 !true
 ```