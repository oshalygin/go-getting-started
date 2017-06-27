package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1337)
	randomNumber := rand.Intn(10)
	fmt.Printf("My favorite number is: %d", randomNumber)
}
