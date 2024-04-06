package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
	fmt.Println("Closed channel")
}

func main() {
	c := make(chan int, 10)
	fibonacci(cap(c), c)

	// The loop for range c receives values from the channel repeatedly until it is closed.
	// If we don't close the channel, the loop will suppose not to stop, the it continue to receive value from channel, which results in deadlock.
	for i := range c {
		fmt.Println(i)
	}
}
