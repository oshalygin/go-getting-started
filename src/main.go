package main

import (
	"fmt"
	"strings"
)

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{hydro, 35, active},
		PowerPlant{hydro, 40, active},
	}

	grid := PowerGrid{1000, plants}

	grid.generatePowerPlantReport()
	grid.generateGridReport()
}

type PlantType string
type PlantStatus string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePowerPlantReport() {
	for index, plant := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", index)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type: ", plant.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", plant.capacity)
		fmt.Printf("%-20s%s\n", "Status: ", plant.status)
		fmt.Printf("\n")
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.00
	for _, plant := range pg.plants {
		if plant.status == active {
			capacity = capacity + plant.capacity
		}
	}

	utilization := capacity / pg.load * 100

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Total Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Grid Load: ", pg.load)
	fmt.Printf("%-20s%.0f %% \n", "Utilization: ", utilization)
}
