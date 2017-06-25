package main

import (
	"fmt"
)

func getCapacity(activePlants []int, plantCapacities []float64) func(float64) float64 {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity = capacity + plantCapacities[plantID]
	}

	return func(gridLoad float64) float64 {
		return getUtilization(capacity, gridLoad)
	}
}

func getUtilization(capacity float64, gridLoad float64) float64 {
	return capacity / gridLoad * 100
}

func printOutput(capacity float64, load float64, utilization float64) {
	fmt.Printf("%-20s%v\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%v\n", "Load: ", load)
	fmt.Printf("%-20s%.1f %%\n", "Utilization: ", utilization)
}

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
		// capacity := calculateCapacity(activePlants, plantCapacities)
		// utilization := getUtilization(capacity, gridLoad)
		capacity := getCapacity(activePlants, plantCapacities)
		utilizationOne := capacity(gridLoad)
		utilizationTwo := capacity(100.0)
		fmt.Printf("%v\n", utilizationOne)
		fmt.Printf("%v\n", utilizationTwo)

		// printOutput(capacity, gridLoad, utilization)

	default:
		fmt.Println("I have no idea what you entered")
	}

}
