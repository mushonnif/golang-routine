package main

import "fmt"

func TestChannel() {
	var message = make(chan string)

	var sayHello = func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		message <- data
	}

	go sayHello("world")
	go sayHello("golang")
	go sayHello("kotlin")

	printMessage(message)
	printMessage(message)
	printMessage(message)
}

func printMessage(message chan string) {
	fmt.Println(<-message)
}
