package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Go Routine (Parallel)")
	var wg sync.WaitGroup
	wg.Add(2)
	author := "Colin Powell"
	quote := "There are no secrets to success. It is the result of preparation, hard work, and learning from failure"
	author2 := "Jim Rohn"
	quote2 := "If you really want to do something, you'll find a way. If you don't, you'll find an excuse"

	go func() {
		printMyQuote(author, quote, 10)
		wg.Done()
	}()

	go func() {
		printMyQuote(author2, quote2, 20)
		wg.Done()
	}()

	fmt.Printf("Done\n")
	wg.Wait()
}

func printMyQuote(author string, quote string, times int) {
	for i := 1; i <= times; i++ {
		fmt.Printf("%s ~ '%s'\n", author, quote)
		time.Sleep(1000 * time.Millisecond)
	}
}
