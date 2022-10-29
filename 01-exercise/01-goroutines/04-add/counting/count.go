package counting

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func sumCon(numbers []int, channel chan int64) {
	var sum int64 = 0
	for _, v := range numbers {
		sum += int64(v)
	}
	
	channel <- sum 
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	
	// Utilize all cores on machine

	// Divide the input into parts
	c1 := make(chan int64)
	c2 := make(chan int64)
	
	

	

	

	// Run computation for each part in seperate goroutine.
	go sumCon(numbers[:len(numbers)/2], c1)
	go sumCon(numbers[len(numbers)/2:], c2)
	

	
	// Add part sum to cummulative sum
	var sum1 int64 = <-c1
	var sum2 int64 = <-c2


	sum = sum1 + sum2
	return sum
}
