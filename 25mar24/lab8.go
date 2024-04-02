package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct to hold podcast data
type Podcast struct {
	Title       string `json:"title"`
	Host        string `json:"host"`
	Episode     int    `json:"episode"`
	Description string `json:"description"`
}

func main() {
	// Take user input
	var title string
	var host string
	var episode int
	var description string

	fmt.Println("Enter podcast title:")
	fmt.Scanln(&title)

	fmt.Println("Enter host name:")
	fmt.Scanln(&host)

	fmt.Println("Enter episode number:")
	fmt.Scanln(&episode)

	fmt.Println("Enter podcast description:")
	fmt.Scanln(&description)

	// Create a Podcast struct with user input
	podcast := Podcast{
		Title:       title,
		Host:        host,
		Episode:     episode,
		Description: description,
	}

	// Marshal Podcast struct into JSON
	jsonData, err := json.Marshal(podcast)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print marshalled JSON
	fmt.Println("Marshalled JSON:")
	fmt.Println(string(jsonData))

	// Unmarshal JSON back into a Podcast struct
	var newPodcast Podcast
	err = json.Unmarshal(jsonData, &newPodcast)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the unmarshalled struct
	fmt.Println("\nUnmarshalled Struct:")
	fmt.Printf("Title: %s\nHost: %s\nEpisode: %d\nDescription: %s\n", newPodcast.Title, newPodcast.Host, newPodcast.Episode, newPodcast.Description)
}
