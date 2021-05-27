package main

import (
	"fmt"
	"sync"
)

// Channel - send data in and out

var wg = sync.WaitGroup{}

func main()  {

	// Recivers & Senders Single channel to communicate their messages across
	ch := make(chan int, 50) // creates 50 internal storage

	wg.Add(2)
	// Recieve only function, channel
	go func (ch <- chan int)  {

		// Reciving Data from the channel

		// Only excepts one data
		i := <-ch
		fmt.Println(i)

		// Accepts all data that are send
		// for j := range ch {
		// 	fmt.Println(j)
		// }

		wg.Done()

		// Select {} - kinda like switch statements but only for channels

	}(ch)

	// Send only channel
	go func (ch chan <- int)  {

		// Send message to channel
		ch <- 42
		close(ch)
		wg.Done()

	}(ch)

	wg.Wait()

}
