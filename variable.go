package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello", "world")
	fmt.Println(string("Hello world"[0]))
	var s string = "Hello world"

	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "world"))
}
