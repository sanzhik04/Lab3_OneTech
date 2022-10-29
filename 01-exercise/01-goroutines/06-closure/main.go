package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output //the output was 4 4 4
	//TODO: fix the issue. //fixed the problem by putting wg.Wait() inside of for-loop

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
		wg.Wait()
	}
}
