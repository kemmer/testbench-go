package main

import (
	"context"
	"fmt"
	"time"
)

const timeout = 4 * time.Second

func main() {
	fmt.Println("=> Context timeout example")
	timeoutExample()

	fmt.Println("=> Context with cancellation example")
	cancelExample()

	fmt.Println("=> Context with deadline example")
	deadlineExample()

	fmt.Println("=> Context value example")
	withValueExample()
}

func cancelExample() {
	ctx, cancel := context.WithCancel(context.Background())

	go performOperationInLoop(ctx, "cancel")

	time.Sleep(timeout)
	cancel()

	time.Sleep(1 * time.Second)
}

func timeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go performOperation()

	select {
	case <-ctx.Done():
		fmt.Println("Context timeout")
	}
}

func deadlineExample() {
	deadline := time.Now().Add(timeout)
	fmt.Println("Deadline is:", deadline)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go performOperationInLoop(ctx, "deadline")

	time.Sleep(time.Second + timeout)
}

func withValueExample() {
	ctx := context.WithValue(context.Background(), "userID", 99)

	performRequest(ctx)
}

func performOperation() {
	select {
	case <-time.After(8 * time.Second):
		fmt.Println("Operation performed after 2 seconds")
	}
}

func performOperationInLoop(ctx context.Context, msgType string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled:", msgType)
			return
		default:
			now := time.Now()
			fmt.Println("Performing operation...", now)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func performRequest(ctx context.Context) {
	userID := ctx.Value("userID").(int)
	fmt.Println("Value available in context:", userID)
}
