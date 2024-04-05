package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 10
	fmt.Println(a)
	fmt.Println(&*&a == &a)
}
