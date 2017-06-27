package main

import "fmt"

func main() {

	firstValue, secondValue := split(50.00)
	fmt.Printf("First Value: %v\n", firstValue)
	fmt.Printf("Second Value: %v\n", secondValue)

}

func split(total float64) (firstValue float64, secondValue float64) {
	firstValue = total / 2
	secondValue = total / 2
	return
}
