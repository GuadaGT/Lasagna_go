package main

import (
	"fmt"
)

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	preparationTime := PreparationTime(numberOfLayers)
	return preparationTime + actualMinutesInOven
}

func main() {

	fmt.Println("Remaining Oven Time:", RemainingOvenTime(30))
	fmt.Println("Preparation Time:", PreparationTime(2))
	fmt.Println("Elapsed Time:", ElapsedTime(3, 20))
}
