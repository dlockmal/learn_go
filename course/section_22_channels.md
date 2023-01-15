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

## Comma ok idiom

## Fan in

## Fan out

## Context