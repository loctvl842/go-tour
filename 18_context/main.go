package main

import (
	"context"
	"fmt"
	"time"
)

func operationWithContext(ctx context.Context, duration time.Duration) {
	// Check if the context is canceled or has reached its deadline
	select {
	case <-ctx.Done():
		fmt.Println("Operation canceled:", ctx.Err())
		return
	case <-time.After(duration):
		fmt.Println("Operation completed successfully")
	}
}

func main() {
	// Create a context with timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Cancel the context to release resources when done

	// Perform an operation with the created context
	operationWithContext(ctx, 1*time.Second)

	// Sleep for 3 seconds to ensure the context's timeout is exceeded
	// time.Sleep(3 * time.Second)

	// Perform another operation with the expired context
	operationWithContext(ctx, 1*time.Second) // This operation will be canceled due to the expired context
}
