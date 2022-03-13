package main

import "fmt"

type User struct {
	Id    int
	Email string
}

func (u *User) String() string {
	return fmt.Sprintf("<%d, %s>", u.Id, u.Email)
}

func main() {
	u := &User{Id: 123, Email: "mail@example.com"}
	fmt.Printf("%s\n", u)
}
