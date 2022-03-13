package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 20)
	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(3 * time.Second)
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}
