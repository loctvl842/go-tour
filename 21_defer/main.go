package main

import (
	"fmt"
	"time"
)

type tradingService struct{}

func (service *tradingService) ListTradeEventBybit() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		go service.ListTradeEventBybit()
	}()

	// Simulate some work
	fmt.Println("Listing trade events...")
	time.Sleep(time.Second)

	// Intentional panic
	panic("Something went wrong!")
}

func main() {
	// Create an instance of the trading service
	service := &tradingService{}

	// Call ListTradeEventBybit in a loop to observe its behavior
	for i := 0; i < 3; i++ {
		service.ListTradeEventBybit()
		time.Sleep(3 * time.Second)
	}

	// Wait for goroutines to finish
	time.Sleep(10 * time.Second)
	fmt.Println("Finished testing.")
}
