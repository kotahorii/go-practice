package main

import (
	"fmt"
	"helloworld/mylib/under"
)

func main() {
	under.Hello()
	person := under.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
	fmt.Println(under.Public)
}
