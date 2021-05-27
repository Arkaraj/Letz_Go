package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Concurrency, Parallelism

func sayHello() {

	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

var wg = sync.WaitGroup{}
var counter int = 0
// OS read-write mutex
// Protect data access
var m = sync.RWMutex{}

func main()  {
	// go sayHello()
	var msg = "Gallo"
	wg.Add(1)
	go func (msg string){
		fmt.Println(msg)
		wg.Done() // decrement by 1
	}(msg)

	msg = "Goodbye"
	
	fmt.Println("Total no. of Threads:", runtime.GOMAXPROCS(-1))

	// Creates 100 threads
	runtime.GOMAXPROCS(100) 
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go Increment()
	}

	wg.Wait()
}

func Increment() {
	counter ++
	m.Unlock()
	wg.Done()
}