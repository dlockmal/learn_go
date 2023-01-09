# Section 20: Concurrency

## Concurrency vs parallelism
- concurrency is built-into go
- parallel is multiple threads/processes at the same design
- concurrency is a design pattern, you can write code that can run in parallel with multiple cores
- concurrency doesn't guarentee parallelism
	- you gotta have more than one cpu
- Talk from Rob Pike on Youtube

## WaitGroup
- godoc.org/sync
- Method sests
	- if the receiver is a pointer it will only work with values that are pointers

----------------------------------------------
Receivers				Values
(t T)									T and *T
(t *T)									*T

```
// func init will run before main
// will run once at the beginning of your script
func init() {

}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	
	wg.Add(1)
	
	// adding 'go' in front of a function will launch 	
	// another go routine. It will be concurrent, but not 	 
	// necessarily parralel
	go foo()
	bar()
	
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
```

## Method sets revisited
wg is type Waitgroup from package sync
```
var wg sync.Waitgroup
```

We are able to use sync.Waitgroup instead of *sync.Waitgroup like the documentation states

**IMPORTANT:** A method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type. 

- Method set determines the interfaces that the type implements

## Documentation
- golang.org -> effective go
- Concurrent programming in many environments is made difficult by the sublteties required to implement correct access to shared variables
- more arden labs diagrams
- synchronization primitives
- mutex and atomic
- go routines are multiplexed onto multiople OS threads so if one should block, such as while waiting for I/O, others can run. Their design hides many of the complexities of thread creation and management.
- Prefix a function or method with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently.
- a function literal can be handy in a goroutine invocation
- https://play.golang.org/p/lBKFKCwrue

## Race condition

## Mutex

## Atomic