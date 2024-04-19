package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func printValues(values ...interface{}) {
	fmt.Println(values...)
}

func main() {
  printValues(1, "Hello", 3.14)
}
