package main

import "fmt"

func main() {
	firstValue := 8
	secondValue := 5
	result := add(firstValue, secondValue)
	fmt.Printf("Result: %v", result)
}
func add(firstValue, secondValue int) int {
	return firstValue + secondValue
}
