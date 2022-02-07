package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 100
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := &Vertex{1, 2, "test"}
	changeVertex2(v)
	fmt.Println(*v)
}
