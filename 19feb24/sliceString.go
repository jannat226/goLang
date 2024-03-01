// Go program to illustrate how to
// create a string from the slice

package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a slice of byte
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	// Creating a string from the slice
	mystring1 := string(myslice1)

	// Displaying the string
	fmt.Println("String 1: ", mystring1)

	// Creating and initializing a slice of rune
	myslice2 := []rune{0x0047, 0x0065, 0x0065, 
							0x006b, 0x0073}

	// Creating a string from the slice
	mystring2 := string(myslice2)

	// Displaying the string
	fmt.Println("String 2: ", mystring2)
}
