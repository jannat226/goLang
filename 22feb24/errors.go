package main

import (
	"errors" //The errors package provides functionality for creating and working with errors
	"fmt"
)

func main() {
	var start, end int
	// Prompt user to enter start and end values
	fmt.Println("Enter the start value:")
	fmt.Scanln(&start)
	fmt.Println("Enter the end value:")
	fmt.Scanln(&end)
	/* This line calls the sumOf function with arguments start and end, and stores the returned values in
	total and err. The := operator is a short variable declaration used to declare and initialize variables.*/
	total, err := sumOf(start, end)
	if err != nil {
		fmt.Println("There was an error:", err)
	} else {
		fmt.Println("The sum from", start, "to", end, "is:", total)
	}
}

/*This line defines a function named sumOf that takes two integers, start and end, as parameters
and returns an integer and an error. The function signature (int, error) indicates that the function
returns an integer and an error object.*/
func sumOf(start int, end int) (int, error) {
	total := 0
	if start > end {
		return 0, errors.New("start is greater than end")
	}
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
