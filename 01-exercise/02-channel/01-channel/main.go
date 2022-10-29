package main

import "fmt"

func main() {

	channel:= make(chan int)
	go func(a, b int) {
		c := a + b
		channel <- c
		close(channel)
	}(1, 2)

	c:= <-channel
	
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", c)
}
