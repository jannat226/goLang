// You are given a list of integers. Your task is to find all the pairs of integers in the list whose sum is equal to a given target value. However, each integer in the list can only be used once in a pair. Write a Go function findPairs that takes in a list of integers numbers, and an integer target, and returns a slice of pairs of integers that sum up to the target. Array elements and the target should be given by the user

// Suppose the array elements are 2, 7, 11, 15, 3, 6, 8, 12    and the target value is  18, the expected output is
// Pairs with sum 18 are: [[7 11] [6 12]]

package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of array you want ")
	fmt.Scanln(&size) // Corrected: scan the size using a pointer

	fmt.Println("Enter the target, i.e. sum of two numbers ")
	var target int
	fmt.Scanln(&target) // Corrected: scan the target using a pointer

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scan(&slice[i])
	}

	fmt.Println("Slice:", slice)

	var matrix [][]int
	// Initialize the 2D slice
	for i := 0; i < size; i++ {
		row := make([]int, 2)
		matrix = append(matrix, row)
	}
	var count int = 0;
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if slice[i]+slice[j] == target {
				fmt.Printf("Sum of pair %d is: %d %d\n", target, slice[i], slice[j])
				matrix[count][0] = slice[i]
				matrix[count][1] = slice[j]
				count ++;
			}
		}
	}

	// Print the 2D slice
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}



