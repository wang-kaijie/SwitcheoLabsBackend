package main

import (
	"fmt"
	"sync"
)

func sum_to_n_a(n int) int {
	/*
		Mathematical Aproach:
		Time Complexity: O(1)
		Space complexity : o(1)
	*/
	return (n * (n + 1)) / 2
}

func sum_to_n_b(n int) int {
	/*
		Iterative Approach with Goroutines and Channels:
		- A new goroutine is launched to compute the sum from 1 to n
		- The computed sum is sent through a channel back to the main function

		Time Complexity: O(n)  -> The function iterates from 1 to n, performing O(n) additions
		Space Complexity: O(1) -> Only a single integer is stored in the channel at any time

		This approach introduces unnecessary concurrency overhead for
		a simple summation task. A direct loop or formula-based approach would
		be more efficient. But it is a unique approach to solve this problem!
	*/

	channel := make(chan int)

	go func() {
		result := 0
		for i := 1; i <= n; i++ {
			result += i
		}
		channel <- result
		close(channel)
	}()

	return <-channel
}

func sum_to_n_c(n int) int {
	/*
		Recursive Approach with Goroutines, Channels, and sync.WaitGroup:
		- A goroutine is spawned for each recursive call
		- The result is passed through a channel back to the caller

		Time Complexity: O(n) -> Recursively processes each number from 1 to n
		Space Complexity: O(n) -> Due to recursion depth and channel usage

		This method is inefficient for large `n` due to excessive goroutine creation
		A simple loop or formula-based approach is preferred
		But it is a unique approach to solve this problem!
	*/

	// Channel to receive final result
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	// Recursive function to compute the sum
	var recursiveSum func(int, chan int)
	recursiveSum = func(num int, subCh chan int) {
		if num == 0 {
			subCh <- 0
			close(subCh)
			return
		}

		// Create a new sub-channel for the next recursive call. add 1 to wg
		nextCh := make(chan int, 1)
		wg.Add(1)

		go func() {
			defer wg.Done()
			recursiveSum(num-1, nextCh)
		}()

		// Receive sum from sub-channel and compute result, close channel afterwards
		subSum := <-nextCh
		subCh <- num + subSum
		close(subCh)
	}

	// Start the recursive process
	wg.Add(1)
	go func() {
		defer wg.Done()
		recursiveSum(n, ch)
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	return <-ch
}

func main() {
	// Change n accordingly
	n := 10
	fmt.Printf("the sum of n: %v", sum_to_n_a(n))
	fmt.Println()
	fmt.Printf("the sum of n: %v", sum_to_n_b(n))
	fmt.Println()
	fmt.Printf("the sum of n: %v", sum_to_n_c(n))
}
