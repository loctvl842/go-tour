package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firtName string
	lastName string
	city     string
	gender   string
	age      int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firtName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
	person1 := Person{firtName: "Samantha", lastName: "Smith", city: "Boston", age: 25, gender: "female"}
	fmt.Println(person1)
  fmt.Println(person1.greet())
}
