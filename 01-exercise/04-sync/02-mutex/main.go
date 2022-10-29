package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex
	var rwm sync.RWMutex


	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		fmt.Println("Someone sent money to balance", balance)
		mu.Unlock()
	}

	read := func() {
		rwm.RLock()
		fmt.Println("Someone is reading balance", balance)
		fmt.Println("Someone's reading session is done", balance)
		rwm.RUnlock()
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
			
		}()

	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			read()
			
		}()

	}






	


	

	//TODO: implement concurrent read.
	// allow multiple reads, writes holds the lock exclusively.

	wg.Wait()
	fmt.Println(balance)
}
