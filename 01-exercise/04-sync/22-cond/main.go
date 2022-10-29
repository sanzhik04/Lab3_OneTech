package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	mutex:= sync.Mutex{}
	condition:=sync.NewCond(&mutex)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		condition.L.Lock()
		for len(sharedRsc) == 0 {
			condition.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		condition.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		condition.L.Lock()
		for len(sharedRsc) == 0 {
			condition.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		condition.L.Unlock()
	}()

	// writes changes to sharedRsc
	condition.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	condition.Broadcast()
	condition.L.Unlock()

	wg.Wait()
}
