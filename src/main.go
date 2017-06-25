package main

import (
	"fmt"
)

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.0

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate power Grid Report")
	fmt.Print("Please choose an option:")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		for index, capacity := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", index, capacity)
		}
	case "2":
		capacity := 0.
		for _, plantID := range activePlants {
			capacity += plantCapacities[plantID]
		}

		utilization := capacity / gridLoad * 100

		// Output
		fmt.Printf("%-20s%v\n", "Capacity: ", capacity)
		fmt.Printf("%-20s%v\n", "Load: ", gridLoad)
		fmt.Printf("%-20s%.1f %%\n", "Utilization: ", utilization)
	default:
		fmt.Println("I have no idea what you entered")
	}
}
