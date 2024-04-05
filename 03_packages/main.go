package main

import (
	"fmt"
	"go-crash-course/03_packages/strutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
  fmt.Println(strutil.Reverse("olleh"))
}
