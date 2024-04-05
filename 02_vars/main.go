package main

import "fmt"

func main() {
	name := "loc"
  age := 20

	fmt.Println("Hello, World!", name)
  fmt.Printf("My name is %v and I'm %v years old\n", name, age)
  // print type of name and age
  fmt.Printf("Type of name is `%T`\n", name)
  fmt.Printf("Type of age is `%T`\n", age)
}
