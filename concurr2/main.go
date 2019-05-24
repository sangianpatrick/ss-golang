package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	defer close(ch)

	rand.Seed(time.Now().UnixNano())

	go func() {
		counting(ch, "AAAA", 5)
	}()

	go func() {
		counting(ch, "BBBB", 5)
	}()

	go func() {
		counting(ch, "CCCC", 5)
	}()

	go func() {
		counting(ch, "DDDD", 5)
	}()

	whoIsFinished := <-ch
	fmt.Printf("%s is finished.\n", whoIsFinished)

}

func counting(ch chan string, person string, multiply int) {
	start := time.Now()
	for i := 1; i <= 10; i++ {
		randTime := int(rand.Intn(1000))
		time.Sleep(time.Duration(randTime) * time.Millisecond)
		result := multiply * i
		fmt.Printf("%s : %d x %d = %d\n", person, i, multiply, result)
	}
	end := time.Since(start)
	fmt.Printf("%s needs %v to finished the task\n", person, end)
	ch <- person
}
