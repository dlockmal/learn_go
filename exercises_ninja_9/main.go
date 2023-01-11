package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func exercise_1() {
	// in addition to the main goroutine, launch two additional goroutines
	// each additional goroutine should print something out
	// use waitgroups to make sure each goroutine finishes before
	// your program exits

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("This is the foo function")
		wg.Done()
	}()

	go func() {
		fmt.Println("This is the bar function")
		wg.Done()
	}()
	wg.Wait()

}

type person struct {
	First string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello", p.First)
}

func saySomething(h human) {
	h.speak()
}

func exercise_2() {
	// More method sets :(
	// create a type person struct
	// attach a method speak to type pointer to a person
	// *person
	// create a type human interface
	// to implicitly implement the interface, a human must have the speak method
	// create a func "saySomething"
	// have it take in a human as a parameter
	// have it call the speak method
	// show the following in your code
	// you CAN pass type *person into saySomething
	// you CANNOT pass type person into saySomething

	p1 := person{"James"}

	saySomething(&p1)
	p1.speak()
}

func exercise_3() {
	// using goroutines, create an incrementer program
	// have a variable to hold the incrementer value
	// launch a bunch of goroutines
	//each goroutine should
	// read the incrementer value
	// store it in a new variable
	// yield the process with runtime.Gosched()
	// increment the new variable
	// write the value in the new variable back to the incrementer variable
	// use waitgroups to wait for all of your goroutines to finish
	// the above will cause a race condition
	// prove that it is a race condition by using the -race flag
	var counter int64
	// counter := 0

	const gs = 100

	var wg sync.WaitGroup
	// var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// mu.Lock()
			// v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			// v++
			// counter = v
			// mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)

	// exercise_4 is to fix the race condition with a mutex
	// exercise_5 is to fix the race condition with atomic

	// this looks ugly with all the comments but it has everything needed
}

func exercise_6() {
	// create a program that prints out your OS and ARCH
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

func main() {
	fmt.Println("Now we wait for Exercise_1 data:")
	exercise_1()
	exercise_2()
	exercise_3() // and exercise_4 and exercise_5
	exercise_6()
}
