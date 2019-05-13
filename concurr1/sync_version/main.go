package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	process1 := addition(20, 30)
	process2 := multiply(process1, 11, 270)
	process3 := division(process2, 788373, 3)
	overallResult := process3
	endTime := time.Since(startTime)
	fmt.Printf("Overall Result: %f\n", overallResult)
	fmt.Printf("Execution Time: %v\n", endTime)
}

func addition(x int, y int) int {
	fmt.Println("Addition: On process")
	result := x + y
	time.Sleep(1 * time.Second)
	fmt.Println("Addition: Done")
	return result
}

func multiply(process1 int, x int, y int) int {
	fmt.Println("Multiply: On process")
	result := x * y
	time.Sleep(1 * time.Second)
	result *= process1
	fmt.Println("Multiply: Done")
	return result
}

func division(process2 int, x int, y int) float64 {
	fmt.Println("Division: On process")
	result := float64(x / y)
	time.Sleep(1 * time.Second)
	result /= float64(process2)
	fmt.Println("Division: Done")
	return result
}
