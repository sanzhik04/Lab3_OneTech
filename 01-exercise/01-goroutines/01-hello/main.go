package main

import (
	"fmt"
	"time"
	"sync"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func apply(s string, f func(s string)){
	f(s)
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(3)
	// Direct call
	fun("direct call")
	
	// TODO: write goroutine with different variants for function call.


	// goroutine function call
	go func(){
		fun("goroutine function call")
		wg.Done()
	}()

	// goroutine with anonymous function
	go func()  {
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine with anonymous function")
			time.Sleep(1 * time.Millisecond)
		}
		wg.Done()
	}()



	// goroutine with function value call
	go func(){
		apply("goroutine with function value call",fun)
		wg.Done()
	}()



	

	

	// wait for goroutines to end
	wg.Wait()
	fmt.Println("done..")
}
