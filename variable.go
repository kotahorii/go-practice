package main

import "fmt"

func foo(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s)

	foo(s...)
}
