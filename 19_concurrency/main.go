package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch)
	}()

  count := 0
	for val, ok := <-ch; count < 9 && ok; val, ok = <-ch {
    fmt.Println("count: ", count)
    // fmt.Println(count)
    count++
		fmt.Println(val)
	}
}
