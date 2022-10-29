// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel

	wg:= new(sync.WaitGroup)
	out := make(chan int)
	output:= func(c <-chan int){
		for value:= range c{
			out<-value
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs{
		go output(c)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := generator(2, 3)

	ch1:= square(in)
	ch2:= square(in)

	

	fmt.Println(<-merge(ch1, ch2))

	// TODO: fan out square stage to run two instances.

	// TODO: fan in the results of square stages.

}
