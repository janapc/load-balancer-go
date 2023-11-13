package main

import (
	"fmt"
)

func workers(workerId int, msgs <-chan string) {
	for msg := range msgs {
		fmt.Printf("[%d] - receive %s \n", workerId, msg)
	}
}

func publish(ch chan<- string) {
	for i := 0; i < 200; i++ {
		ch <- fmt.Sprintf("message number #%d", i)
	}
}

func main() {
	ch := make(chan string)
	totalWorkers := 20

	for i := 0; i < totalWorkers; i++ {
		go workers(i, ch)
	}
	publish(ch)
}
