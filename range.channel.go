package main

import "fmt"

func sendMessage(message chan<- string) {
	for i := 0; i < 10; i++ {
		fmt.Println("sending: ", i)
		message <- "message: " + fmt.Sprint(i)
	}

	close(message)
}

func TestRangeChannel() {
	var message = make(chan string)
	go sendMessage(message)

	for msg := range message {
		fmt.Println(msg)
	}
}
