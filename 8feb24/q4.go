// Q4
// You are tasked with creating a grade calculator program in GoLang. 
// The program should prompt the user to input their scores in three subjects: Math, Science, and Science. 
// Based on these scores, calculate the average grade (assuming each subject is equally weighted) and display 
//the corresponding letter grade (A, B, C, D, or F) using control flow.

package main

import (
	"fmt"
)

func main() {
	// Prompting the user to input scores for three subjects
	var mathScore, scienceScore, englishScore float64
	fmt.Println("Enter your score for Math:")
	fmt.Scanln(&mathScore)
	fmt.Println("Enter your score for Science:")
	fmt.Scanln(&scienceScore)
	fmt.Println("Enter your score for English:")
	fmt.Scanln(&englishScore)

	// Checking if the entered marks are within the valid range
	if mathScore < 0 || mathScore > 100 || scienceScore < 0 || scienceScore > 100 || englishScore < 0 || englishScore > 100 {
		fmt.Println("Invalid input. Scores should be between 0 and 100.")
		return
	}

	// Calculating the average grade
	average := (mathScore + scienceScore + englishScore) / 3

	// Determining the letter grade based on the average
	var letterGrade string
	if average >= 90 {
		letterGrade = "A"
	} else if average >= 80 {
		letterGrade = "B"
	} else if average >= 70 {
		letterGrade = "C"
	} else if average >= 60 {
		letterGrade = "D"
	} else {
		letterGrade = "F"
	}

	// Displaying the average grade
	fmt.Printf("Your average grade is %.2f, which corresponds to letter grade: %s\n", average, letterGrade)
}
