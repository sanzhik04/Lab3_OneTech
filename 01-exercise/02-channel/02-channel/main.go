package main



import "fmt"

func main() {

	channel:= make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			// TODO: send iterator over channel
			channel <- i
		}
		close(channel)
	}()

	// TODO: range over channel to recv values

	for i := range channel {
		fmt.Println(i)
	}
	


	

}
