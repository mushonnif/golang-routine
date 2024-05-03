package main

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func TestWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doPrint(&wg, fmt.Sprintf("message: %d", i))
	}

	wg.Wait()
}
