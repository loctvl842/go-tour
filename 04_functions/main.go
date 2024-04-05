package main

import (
	"fmt"
	"go-crash-course/04_functions/helper"
)

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(helper.Greeting("Loc"))
	fmt.Println(sum(2, 1))
}
