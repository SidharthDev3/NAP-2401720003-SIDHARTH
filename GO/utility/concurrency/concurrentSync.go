package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func ConcurrentSync() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		printMessage("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		printMessage("Goroutine 2")
	}()

	printMessage("Hi!!")

	wg.Wait()
}
