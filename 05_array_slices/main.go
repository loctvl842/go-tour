package main

import "fmt"

func main() {
	var fruitArr [2]string
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"

	// Declare and assign
	fruitArr2 := []string{"apple", "orange", "grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)
  fmt.Println(len(fruitArr2))
  fmt.Println(fruitArr2[1:2])
}
