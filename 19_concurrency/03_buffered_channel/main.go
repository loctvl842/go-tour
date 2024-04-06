package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // This will cause a deadlock
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch) // This will cause a deadlock
}
