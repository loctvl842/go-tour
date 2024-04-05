package main

import "fmt"

func printValue(v interface{}) {
	fmt.Println("Value:", v)
}

func main() {
	printValue(42)             // Value: 42
	printValue("hello")        // Value: hello
	printValue(3.14)           // Value: 3.14
	printValue([]int{1, 2, 3}) // Value: [1 2 3]
}
