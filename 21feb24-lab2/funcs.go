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
	Name, HostName, Speaker string
}

// Implementing the Streamer interface for Podcast
func (p Podcast) Stream() {
	fmt.Printf("Streaming podcast '%s' hosted by %s, with speaker %s\n", p.Name, p.HostName, p.Speaker)
}

func main() {
	// Print welcome message and program overview
	fmt.Println("Welcome to the Podcast Streaming System.")
	fmt.Println("This program allows you to input details about podcasts, including the podcast name, host name, and speaker name.")
	
	// Prompt the user to enter the number of podcasts
	fmt.Println("Enter the number of podcasts you want to input details for:")
	var count int
	fmt.Scanln(&count)

	// Slice to store podcast details
	podcasts := make([]Podcast, count)

	// Maps to store counts of hosts and speakers
	hostCount := make(map[string]int)
	speakerCount := make(map[string]int)

	// Input podcast details
	for i := 0; i < count; i++ {
		fmt.Printf("\nEnter details for Podcast %d\n", i+1)

		// Read podcast name
		name := readPodcastName(podcasts)

		// Read host and speaker names
		fmt.Print("Enter host name: ")
		host := getInput()
		fmt.Print("Enter speaker name: ")
		speaker := getInput()

		// Update host and speaker counts
		hostCount[host]++
		speakerCount[speaker]++

		// Store podcast details in slice
		podcasts[i] = Podcast{Name: name, HostName: host, Speaker: speaker}
	}

	// Print overview of entered podcast details
	fmt.Println("\nOverview of entered podcast details:")
	for i, podcast := range podcasts {
		fmt.Printf("\nPodcast %d:\nName: %s\nHost: %s\nSpeaker: %s\n", i+1, podcast.Name, podcast.HostName, podcast.Speaker)

		// Stream the podcast
		podcast.Stream()
	}

	// Print counts of hosts
	fmt.Println("\nCounts of hosts:")
	for host, count := range hostCount {
		fmt.Printf("Host: %s, Podcast Count: %d\n", host, count)
	}

	// Print counts of speakers
	fmt.Println("\nCounts of speakers:")
	for speaker, count := range speakerCount {
		fmt.Printf("Speaker: %s, Podcast Count: %d\n", speaker, count)
	}
}

// Function to read podcast name and validate if it already exists
func readPodcastName(podcasts []Podcast) string {
	var name string
	for {
		fmt.Print("Enter podcast name: ")
		name = getInput()

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
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
