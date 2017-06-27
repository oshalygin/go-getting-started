package main

import "fmt"

func main() {

	firstValue, secondValue := swap(1, 2)

	fmt.Printf("firstValue: %v\n", firstValue)
	fmt.Printf("secondValue: %v\n", secondValue)
}

func swap(firstValue int, secondValue int) (int, int) {
	return secondValue, firstValue
}
