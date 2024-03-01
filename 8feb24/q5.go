package main

import (
	"fmt"
)

func main() {
	var numSets int
	fmt.Println("Enter the number of sets of data:")
	fmt.Scanln(&numSets)

	principalAmounts := make([]float64, numSets)
	annualInterestRates := make([]float64, numSets)
	numberOfYears := make([]int, numSets)
	simpleInterests := make([]float64, numSets)

	// Input data for each set
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nEnter data for set %d:\n", i+1)

		fmt.Print("Principal Amount: ")
		fmt.Scanln(&principalAmounts[i])

		fmt.Print("Annual Interest Rate: ")
		fmt.Scanln(&annualInterestRates[i])

		fmt.Print("Number of Years: ")
		fmt.Scanln(&numberOfYears[i])

		// Calculate simple interest for each set
		simpleInterests[i] = (principalAmounts[i] * annualInterestRates[i] * float64(numberOfYears[i])) / 100
	}

	// Display the results for each set
	fmt.Println("\nResults:")
	for i := 0; i < numSets; i++ {
		fmt.Printf("Set %d:\n", i+1)
		fmt.Printf("Principal Amount: %.2f\n", principalAmounts[i])
		fmt.Printf("Annual Interest Rate: %.2f%%\n", annualInterestRates[i])
		fmt.Printf("Number of Years: %d\n", numberOfYears[i])
		fmt.Printf("Simple Interest: %.2f\n", simpleInterests[i])
		fmt.Println("----------------------")
	}
}
