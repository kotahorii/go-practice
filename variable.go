package main

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(1, 2)
	fmt.Println(r1, r2)

	fmt.Println(cal(100, 2))

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
