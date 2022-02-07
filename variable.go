package main

import "fmt"

func main() {
	n := 100
	fmt.Println(&n)

	p := &n
	fmt.Println(*p)
}
