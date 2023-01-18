# Section 22: Channels

## Understanding channels
- a better way of writing concurrent code and writing better code
- channels block
- ==when you send and receive on a channel, they have to happen at the same time, otherwise there's a deadlock==
	- **Understanding Channels Block is imperitive**
- There are two easy ways to do this

Method #1 (synchronized)
```
func main() {
	c := make(chan int)
	
	go func(){
		c <- 42
	}()
	
	fmt.Println(<-c)
}
```

Method #2 (Buffer)
```
func main() {
	// can change the buffer to 2 to work
	c := make(chan int, 1)
	c <- 42
	// if we add this below we'll get a deadlock
	// because our buffer is one and we're trying
	// to add 43 to a buffer of 1 with 42 already 
	// there
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
```

- buffers are a slippery slope, they're okay if they're done right. 
	- You don't know when you'll reach your buffer threashold
- There's a time and place for buffer channels, but in general probably best to stay away
- do not communicate by sharing memory, instead, share memory by communicating

- Go Cones... "Rob Pike Go Sayings" / Go Proverbs
- This is the last type in go... we've learned them all... holy shit


## Directional channels
- You can mark channels as receive only
	- so your only pulling values off and not putting data on the channel
```
func main() {
	// this is a channel that can only send
	c := make(chan<- int, 2)

	// this is a channel that can only receive
	c := make(<-chan int, 2)
	
	// this will fail and succeed depending on your channel type
	c <- 42
	c <- 43
	
	fmt.Println(<-chan)
	fmt.Printf("%T\n", c)
	 
}
```
- channel type is read from left to right
- you can take a bidirectional channel and make it more specific (send / receive) but you can't take a specific channel and make it more generic


## Using channels
```
func main() {

c := make(chan int)

	//send
	go foo(c)
	
	//receive
	bar(c)
	
	fmt.Println("about to exit")
}

	//send only channel
	func foo(c chan<- int){
		c <- 42
	}
	
	//receive
	func bar(c <-chan int){
		fmt.Println(<-c)
	}
```

## Range
- A nice design pattern
```
func main() {

	c := make(chan int)

	//send
	gofunc(){
	for i := 0; i < 100; i++{
			c <- i
		}
		close(c)
	}()
	
	//receive
	for v := range c {
		fmt.Println(v)
	}
	
	fmt.Println("about to exit")
}
```

## Select
- channels block
- Can use range or separate go routines to prevent
```
func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	
	// send
	go send(eve, odd, quit)
	
	// receive
	receive(eve, odd, quit)
	
	fmt.Println("about to exit")
}

func receive (e, o, q <- chan int) {
	// infinite loop until we get out of it
	for {
		select {
		case v := <- e:
			fmt.Println("from the eve channel:", v)
		case v := <- o:
			fmt.Println("from the odd channel:", v)
		case v := <- q:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int){
	for i := 0; i < 100, i++{
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
```

## Comma ok idiom with select statements
- comma ok idiom with channels
- there are other examples in the documentation
```
func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	
	v, ok := <-c
	
	fmt.Println(v, ok)
	fmt.Println(<-c)
}
```

## Fan in
- fanned from multiple channels onto a single channel
- nice design pattern, but a little complex
```
func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)
	
	go send(even, odd)
	
	go receive(even, odd, fanin)
	
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}

// send channel
func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()
	
	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
```

## Fan out
- taking a piece of work, and instead of doing the work sequential(serial), do it parallel
- You can also throttle this work
```
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	
	go populate(c1)
	
	go fanOutIn(c1, c2)
	
	for v := range c2 {
		fmt.Println(v)
	}
	
	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i>
	}
	close(c)
}

func canOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
```

for throttling the code would look like this

```
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

```

## Context
- a tool we can use for our concurrent design process launches other go routines, that when we cancel the top process all the goroutines cancel
- We don't want to leak goroutines
- passing around variables related to request
- we need to know about as we get deeper into go but it's more of an advanced topic
- the go blog by sameer 'Go Concurrency Patterns: Context'
- good examples of request scoped data include user IDs, extracted headers, authentication tokens, or session ids
- How did he know the type Context was a struct by looking at the docs? It's got to be in there somewhere
- takaway general idea of contexts
```
func main() {
	ctx := context.Background()
	
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("context type:\t%T"\n, ctx)
	fmt.Println("-----------")
	
	// pass in a parent context
	ctx, _ = context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("cancel:\t\t", cancel)
	fmt.Println("context type:\t%T"\n, ctx)
	fmt.Println("cancel type:\t%T\n", cancel)
	fmt.Println("-----------")
	
	cancel()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("cancel:\t\t", cancel)
	fmt.Println("context type:\t%T"\n, ctx)
	fmt.Println("cancel type:\t%T\n", cancel
	fmt.Println("-----------")
}

```

and here's a legit example
```
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumBoroutine())
	
	go func() {
		n := 0
		for {
			select {
				case <-ctx.Done():
					return
				default:
					n++
					time.Sleep(time.Millisecond * 200)
					fmt.Println("working", n)
			}
		}
	}()
	
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())
	
	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")
	
	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}
```

- when you create a service, you want to be very careful with goroutines and not leak any of them