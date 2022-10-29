package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		var n int = 0
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case out <- n:
					n++
				}
			}

		}()

		return out
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for value := range generator(ctx) {
		fmt.Println(value)
		if value == 5 {
			cancel()
			break
		}
	}

}
