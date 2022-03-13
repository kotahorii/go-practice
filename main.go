package main

import "fmt"

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

type T struct {
	Id   int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func PrintStr(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	PrintStr(&T{Id: 3, Name: "Takashi"})
}
