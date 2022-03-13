package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
}

func main() {
	src := `
		{
			"Id":1,
			"Name":"Takashi",
			"Email":"example.com",
			"Created":"2022-03-13T15:26:11.178205+09:00"
		}
		`
	u := new(User)
	err := json.Unmarshal([]byte(src), u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
}
