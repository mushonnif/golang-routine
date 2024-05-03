package main

import "fmt"

func print(till int, name string) {
	for i := 0; i < till; i++ {
		println(name, i)
	}
}
func TestRoutine() {
	go print(10, "hello")
	print(10, "world")

	var input string
	fmt.Scan(&input)
}
