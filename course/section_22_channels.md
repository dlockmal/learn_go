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

## Using channels

## Range

## Select

## Comma ok idiom

## Fan in

## Fan out

## Context