package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Symlink("foo.txt", "bar.txt")
	if err != nil {
		log.Fatal(err)
	}
	path, _ := os.Readlink("bar.txt")
	fmt.Println(path)
}
