package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan float64)

	go func() {
		addition(ch1, 20, 30)
		close(ch1)
	}()

	go func() {
		multiply(ch1, ch2, 11, 270)
		close(ch2)
	}()

	go func() {
		division(ch2, ch3, 788373, 3)
		close(ch3)
	}()

	overallResult := <-ch3
	endTime := time.Since(startTime)
	fmt.Printf("Overall Result: %f\n", overallResult)
	fmt.Printf("Execution time: %v\n", endTime)

}

func addition(ch1 chan int, x int, y int) {
	fmt.Println("Addition: On process")
	result := x + y
	time.Sleep(1 * time.Second)
	fmt.Println("Addition: Done")
	ch1 <- result
}

func multiply(ch1 chan int, ch2 chan int, x int, y int) {
	fmt.Println("Multiply: On process")
	result := x * y
	time.Sleep(1 * time.Second)
	fmt.Println("Multiply: Waiting for result of addition")
	result *= <-ch1
	time.Sleep(1 * time.Second)
	fmt.Println("Multiply: Done")
	ch2 <- result
}

func division(ch2 chan int, ch3 chan float64, x int, y int) {
	fmt.Println("Division: On process")
	result := float64(x / y)
	time.Sleep(1 * time.Second)
	fmt.Println("Division: Waiting for result of multiply")
	result /= float64(<-ch2)
	time.Sleep(1 * time.Second)
	fmt.Println("Division: Done")
	ch3 <- result
}
