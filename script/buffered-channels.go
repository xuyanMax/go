package main

import "fmt"
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

// ch := make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

// Modify the example to overfill the buffer and see what happens.
func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 10
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
