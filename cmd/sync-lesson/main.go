package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(400 * time.Millisecond)
	ticker2 := time.NewTicker(20 * time.Millisecond)
	done := make(chan bool)
	retriesCount := make(map[string]int)
	mutex := &sync.RWMutex{}

	go tick(done, ticker, retriesCount, mutex)
	go tick(done, ticker2, retriesCount, mutex)

	time.Sleep(400*10*time.Millisecond + 2*time.Millisecond)
	ticker.Stop()

	done <- true

	for k, v := range retriesCount {
		fmt.Printf("%v: %v", k, v)
	}
}

func tick(done chan bool, ticker *time.Ticker, rc map[string]int, mutex *sync.RWMutex) {
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("tick: ", t)

			mutex.Lock()
			rc["a"] += 1
			mutex.Unlock()
		}
	}
}
