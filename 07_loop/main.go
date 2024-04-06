package main

import "fmt"

func NormalLoop() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func InfiniteLoop() {
  for {
    fmt.Println("Infinite Loop")
  }
}

func main() {
  InfiniteLoop()
}
