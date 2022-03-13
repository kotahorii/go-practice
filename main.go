package main

import (
	"fmt"
	"log"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func NumIsGreaterThanZero(num int) (int, error) {
	if num >= 0 {
		return num, nil
	}
	return 0, &MyError{Message: "エラーが発生しました"}
}

func main() {
	num, err := NumIsGreaterThanZero(-10)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(num)
}
