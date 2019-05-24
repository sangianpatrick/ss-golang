package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	ch <- "asdf"
	fmt.Println(<-ch)
}
