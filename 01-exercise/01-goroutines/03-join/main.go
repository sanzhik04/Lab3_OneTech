package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	wg := new(sync.WaitGroup)
	wg.Add(1)
	var data int

	go func() {
		data++
		defer wg.Done()
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)

	fmt.Println("Done..")
}
