package main

import (
	"fmt"
)

func main() {
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	capacity := plantCapacities[0] + plantCapacities[1] + plantCapacities[2]

	fmt.Println(plantCapacities)
	fmt.Println(capacity)

	gridLoad := 75.

	utilization := gridLoad / capacity

	fmt.Printf("Capacity: %v\n", capacity)
	fmt.Printf("Utilization: %.2f\n", utilization*100)

}
