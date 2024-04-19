package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Create a new FlagSet with custom usage message
	fs := flag.NewFlagSet("example", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", fs.Name())
		fmt.Fprintf(os.Stderr, "  -name string\n\tThe name to greet\n")
		fmt.Fprintf(os.Stderr, "  -times int\n\tNumber of times to greet\n")
	}

	// Define flags using the FlagSet
	name := fs.String("name", "World", "The name to greet")
	times := fs.Int("times", 1, "Number of times to greet")

	// Parse command-line flags
	fs.Parse(os.Args[1:])

	// Print greetings
	for i := 0; i < *times; i++ {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
