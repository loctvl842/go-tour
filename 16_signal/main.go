package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Create a channel to receive OS signals
	signalChan := make(chan os.Signal, 1)

	// Notify the signalChan channel of the specified signals
	// signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
  go func() {
    signalChan <- syscall.SIGHUP
  }()

	// Block until a signal is received
	sig := <-signalChan

	fmt.Printf("Received signal: %v\n", sig)
}
