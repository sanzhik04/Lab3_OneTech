package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {

			defer close(ch)
			select {
			case <-time.After(50 * time.Millisecond):
				// Simulate work.
				ch <- data{"123"}
				fmt.Println("Successful sending")
			case <-ctx.Done():
				// Simulate work.
				fmt.Println(ctx.Err())
				return
				// Report result.
				

			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	deadline := time.Now().Add(50* time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	ch := compute(ctx)
	d := <-ch
	fmt.Printf("work complete: %s\n", d)

}
