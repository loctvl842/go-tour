package helper

import "strings"

func Greeting(name string) string {
  var uppercaseName string
  if len(name) > 0 {
    uppercaseName = strings.ToUpper(name[:1]) + name[1:]
  }
  return "Hello " + uppercaseName
}

// Write fibonacci function
