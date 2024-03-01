// Q3
// You're developing a weather monitoring system in GoLang for a research institute. 
// The system needs to store data about temperature, humidity, and wind speed. 
// Define variables to hold these measurements for a specific location at a given point in time.
//  Ensure you choose suitable data types for representing numerical measurements accurately.

package main

import "fmt"

type Weather struct {
	Temp      float32
	Humidity  float32
	WindSpeed float32
}

func main() {
	var noOfLoc int
	fmt.Println("Enter the number of Locations you want")
	fmt.Scanln(&noOfLoc)

	// Create a slice to store locations
	locations := make([]Weather, noOfLoc)
	mp := make(map[Weather]string)

	// Taking input from user
	for i := 0; i < noOfLoc; i++ {
		var location string
		fmt.Printf("Enter details for location %d:\n", i+1)
		fmt.Print("Location: ")
		fmt.Scanln(&location)

		var temp, humidity, wind float32
		fmt.Print("Temp: ")
		fmt.Scanln(&temp)
		fmt.Print("Humidity: ")
		fmt.Scanln(&humidity)
		fmt.Print("WindSpeed: ")
		fmt.Scanln(&wind)

		locations[i] = Weather{Temp: temp, Humidity: humidity, WindSpeed: wind}
		mp[locations[i]] = location
	}

	// Displaying the map
	fmt.Println("Weather details for each location:")
	for _, loc := range locations {
		fmt.Printf("Location: %s\n", mp[loc])
		fmt.Printf("Temp: %.2f\n", loc.Temp)
		fmt.Printf("Humidity: %.2f\n", loc.Humidity)
		fmt.Printf("WindSpeed: %.2f\n", loc.WindSpeed)
		fmt.Println("--------------------------")
	}
}
