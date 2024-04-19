package main

import "fmt"

func main() {
	dict := map[string]interface{}{}
	dict["name"] = "loc"
	dict["age"] = 20
	is, ok := dict["age"]
	fmt.Println(dict["name"])
	fmt.Println(is, ok)
}
