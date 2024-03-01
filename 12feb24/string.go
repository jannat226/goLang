package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("String Functions as follow: ")
	reader := bufio.NewReader(os.Stdin)

	// 1. Length of a string
	fmt.Println("Input a string to perform the following operations")
	fmt.Println("Enter your choice based on the menu \n1. Length of string \n2. Concatenation \n3. ToUpper \n4. ToLower \n5. Split")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Length")
		str, _ := reader.ReadString('\n')
		fmt.Println("1. Length of string:", len(str))
	case 2:
		fmt.Println("Concatenation")
		// Concatenation
		str1, _ := reader.ReadString('\n')
		str2, _ := reader.ReadString('\n')
		concatenated := strings.TrimSpace(str1) + " " + strings.TrimSpace(str2)
		fmt.Println("The concatenated string is", concatenated)
	case 3:
		fmt.Println("ToUpper")
		// Conversion to uppercase
		lwStr, _ := reader.ReadString('\n')
		upStr := strings.ToUpper(strings.TrimSpace(lwStr))
		fmt.Println("Changed to upper case:", upStr)
	case 4:
		fmt.Println("ToLower")
		// Conversion to lower
		upperStr, _ := reader.ReadString('\n')
		lowerStr := strings.ToLower(strings.TrimSpace(upperStr))
		fmt.Println("Changed to lower case:", lowerStr)
	case 5:
		fmt.Println("Split")
		// Split
		fmt.Println("Enter a string with comma as a delimiter:")
		fruits, _ := reader.ReadString('\n')
		splitted := strings.Split(strings.TrimSpace(fruits), ",")
		fmt.Println(splitted)
		for _, fruit := range splitted {
			fmt.Println("-", fruit)
		}
	default:
		fmt.Println("Invalid")
	}
}
