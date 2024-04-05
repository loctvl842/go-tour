package main

import "fmt"

func main() {
	emails := make(map[string]string)
	emails["Loc"] = "onno@yopmail.com"
  emails["Van"] = "xiumui2012002@yopmail.com"

  fmt.Println(emails)
  fmt.Println(len(emails))
  fmt.Println(emails["Loc"])

  // Declare map and add key-value
  emails2 := map[string]string{"Loc": "onno@yopmail.com"}
  fmt.Println(emails2)
}
