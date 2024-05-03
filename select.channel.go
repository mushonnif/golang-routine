package main

import "fmt"

func getAverage(numbers []int, c chan float64) {
	var sum = 0
	for _, value := range numbers {
		sum += value
	}
	c <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, c chan int) {
	var max = numbers[0]
	for _, value := range numbers {
		if max < value {
			max = value
		}
	}
	c <- max
}

func TestSelectChannel() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("numbers: ", numbers)

	var c1 = make(chan float64)
	var c2 = make(chan int)
	go getAverage(numbers, c1)
	go getMax(numbers, c2)

	for i := 0; i < 2; i++ {
		select {
		case average := <-c1:
			fmt.Println("average: ", average)
		case max := <-c2:
			fmt.Println("max: ", max)
		}
	}
}
