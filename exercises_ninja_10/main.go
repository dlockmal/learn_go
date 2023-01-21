package main

import "fmt"

func main() {
	exercise_1()
	exercise_2_1()
	exercise_2_2()
	exercise_3()
	exercise_4()
	exercise_5()
	exercise_6()
	exercise_7()
}

func exercise_1() {
	// we're going to start with some code that we need to get working
	// using func literal, aka, anonymous self-executing func
	// a buffered channel
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("Exercise 1 Output: ")
	fmt.Println(<-c)
}

func exercise_2_1() {
	// more get this code running
	// no hints this time
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func exercise_2_2() {
	// these were worthless I just modified the channels from send or receive only to not specific
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func exercise_3() {
	// start with this code
	// pull the values off the channel using a for range loop
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func exercise_4() {
	// start with the code given
	// pull the values off the channel using a select statement
	q := make(chan int)
	c := gen4(q)

	receive4(c, q)

	fmt.Println("about to exit")
}

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive4(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("from the channel:", v)
		case <-q:
			return
		}
	}
}

func exercise_5() {
	// show the comma ok idiom starting with the code
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func exercise_6() {
	// write a program that puts 100 numbers to a channel
	// pull the numbers off the channel and print them
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exercise 6: About to exit")
}

func exercise_7() {
	// write a program that launches 10 goroutines
	// each goroutine adds 10 numbers to a channel
	// pull the numbers off the channel and print them
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for x := 0; x < 10; x++ {
				c <- x
			}
		}()
	}

	for v := 0; v < 100; v++ {
		fmt.Println(v, <-c)
	}
	fmt.Println("Exercise 7 about to close")
}
