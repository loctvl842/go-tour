package main

import "fmt"

func main() {
  x := 10
  y := 5

  if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
  } else {
    fmt.Printf("%d is less than %d\n", y, x)
  }
}
