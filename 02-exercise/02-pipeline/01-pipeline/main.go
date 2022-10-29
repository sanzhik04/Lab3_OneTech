package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int{
	output:= make(chan int)
	
	go func(){
		for _,value := range nums{
			output <- value
		}
		close(output)
    }()

	return output


}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(input <-chan int) <-chan int{
	output:= make(chan int)
	go func(){
		for value:= range input{
			output <- value*value
		}

		close(output)
	}()

	return output
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	channel1:= square(generator(1,2,3,4,5))
	channel2:=square(square(generator(1,2,3,4,5)))

	for value:= range channel1{
		fmt.Println(value)
	}

	for value :=range channel2{
		fmt.Println(value)
	}


	

	

}
