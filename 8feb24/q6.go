package main

import (
	"fmt"
)

// Function to calculate factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	const numPodcasts = 3 // Constant for the number of podcasts
	var podcastIDs [numPodcasts]int

	// Taking input for podcast IDs
	for i := 0; i < numPodcasts; i++ {
		fmt.Printf("Enter the ID for Podcast %d: ", i+1)
		fmt.Scanln(&podcastIDs[i])
	}

	// Calculate factorial for each podcast ID and display the result
	fmt.Println("\nFactorial of podcast IDs:")
	for _, id := range podcastIDs {
		fact := factorial(id)
		fmt.Printf("Podcast ID: %d, Factorial: %d\n", id, fact)
	}
}
