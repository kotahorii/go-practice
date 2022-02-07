package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	n = append(n, 100)
	fmt.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)
}
