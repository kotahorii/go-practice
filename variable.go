package main

import "fmt"

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	c1 := circleArea(3.14)
	fmt.Println(c1(3))
}
