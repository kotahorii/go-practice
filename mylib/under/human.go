package under

import "fmt"

var Public string = "Public"
var private string = "private"

type Person struct {
	Name string
	Age  int
}

func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}

func Say() {
	fmt.Println("Human!")
}
