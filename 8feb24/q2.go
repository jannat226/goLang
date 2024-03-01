// Q2
// You're tasked with building a student information system in GoLang for a school. 
// Each student record needs to store details such as student ID, name, age, and grade.
//  Define variables to store the information of a single student and assign values accordingly. 
// Pay attention to selecting appropriate data types to represent each piece of information.
// package main

package main

import "fmt"

type Student struct {
	Id, Name, Grade string
	Age             int
}

func main() {
	var noOfStudent int
	var name, grade, id string
	var age int

	fmt.Println("Enter the number of students you want to enter:")
	fmt.Scanln(&noOfStudent)

	students := make([]Student, noOfStudent)

	for i := 0; i < noOfStudent; i++ {
		fmt.Println("Enter the details of student", i+1)
		fmt.Println("Student Id:")
		fmt.Scanln(&id)
		fmt.Println("Student Name:")
		fmt.Scanln(&name)
		fmt.Println("Student Grade:")
		fmt.Scanln(&grade)
		fmt.Println("Student Age:")
		fmt.Scanln(&age)
		students[i] = Student{Id: id, Name: name, Grade: grade, Age: age}
	}

	// Print out the students
	fmt.Println("\nStudents in list are:")
	for i, student := range students {
		fmt.Printf("Student id %d:\n", i+1)
		fmt.Printf("Name: %s\n", student.Name)
		fmt.Printf("Grade: %s\n", student.Grade)
		fmt.Printf("Age: %d\n", student.Age)
	}
}