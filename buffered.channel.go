package main

import (
	"fmt"
	"time"
)

func TestBufferedChannel() {
	var message = make(chan int, 2)

	go func() {
		for {
			var data = <-message
			fmt.Printf("received: %d\n", data)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("sending: %d\n", i)
		message <- i
	}

	time.Sleep(1 * time.Second)
}
