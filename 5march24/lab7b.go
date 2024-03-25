package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Streamer interface to represent entities that stream content
type Streamer interface {
	Stream()
}

// Podcast struct to store podcast details
type Podcast struct {
	Name     string
	HostName string
	Speakers []string
}

// Implementing the Streamer interface for Podcast
func (p Podcast) Stream() {
	fmt.Printf("Streaming podcast %s hosted by %s, with speakers %s\n", p.Name, p.HostName, strings.Join(p.Speakers, ", "))
}

func main() {
	fmt.Println("Welcome to the Podcast Streaming System.")
	fmt.Println("This program allows you to input details about podcasts, including the podcast name, host name, and speaker names.")

	// Slice to store podcast details
	var podcasts []Podcast

	// Maps to store counts of hosts and speakers
	hostCount := make(map[string]int)
	speakerCount := make(map[string]int)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Enter podcast details")
		fmt.Println("2. View counts of hosts and speakers")
		fmt.Println("3. View average number of speakers per podcast")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nEnter details for a new podcast:")
			// Read podcast name
			name := readPodcastName(podcasts)

			// Read host name
			fmt.Print("Enter host name: ")
			host, err := getInput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Read speaker names
			fmt.Print("Enter speaker names (comma-separated): ")
			speakerInput, err := getInput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			speakers := strings.Split(speakerInput, ",")

			// Update host and speaker counts
			updateCounts(&hostCount, &speakerCount, host, speakers)

			// Store podcast details in slice
			podcast := Podcast{Name: name, HostName: host, Speakers: speakers}
			podcasts = append(podcasts, podcast)

		case 2:
			// Print counts of hosts
			fmt.Println("\nCounts of hosts:")
			printCounts(hostCount)

			// Print counts of speakers
			fmt.Println("\nCounts of speakers:")
			printCounts(speakerCount)

		case 3:
			// Calculate average number of speakers per podcast
			if len(podcasts) == 0 {
				fmt.Println("No podcasts entered yet.")
			} else {
				avgSpeakers := calculateAverageSpeakers(podcasts)
				fmt.Printf("\nAverage number of speakers per podcast: %d\n", avgSpeakers)
			}

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}

// Function to read podcast name and validate if it already exists
func readPodcastName(podcasts []Podcast) string {
	var name string
	var err error
	for {
		fmt.Print("Enter podcast name: ")
		name, err = getInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Check if the podcast name already exists
		exists := false
		for _, podcast := range podcasts {
			if podcast.Name == name {
				exists = true
				fmt.Println("Podcast with the same name already exists. Please enter a different name.")
				break
			}
		}

		if !exists {
			break
		}
	}
	return name
}

// Function to read input
func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// Function to update host and speaker counts
func updateCounts(hostCount, speakerCount *map[string]int, host string, speakers []string) {
	(*hostCount)[host]++
	for _, speaker := range speakers {
		(*speakerCount)[speaker]++
	}
}

// Function to print counts
func printCounts(counts map[string]int) {
	for key, value := range counts {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// Function to calculate the average number of speakers per podcast
func calculateAverageSpeakers(podcasts []Podcast) int {
	totalSpeakers := 0
	for _, podcast := range podcasts {
		totalSpeakers += len(podcast.Speakers)
	}
	return totalSpeakers / len(podcasts)
}

