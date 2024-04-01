package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// Podcast struct to store podcast details
type Podcast struct {
	Name     string   `json:"name"`
	HostName string   `json:"hostName"`
	Speakers []string `json:"speakers"`
}

// Streamer interface to represent entities that stream content
type Streamer interface {
	Stream()
}

// Implementing the Streamer interface for Podcast
func (p Podcast) Stream() {
	fmt.Printf("Streaming podcast %s hosted by %s, with speakers %s\n", p.Name, p.HostName, strings.Join(p.Speakers, ", "))
}

// MarshalJSON encodes the Podcast struct to JSON
func (p Podcast) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name     string   `json:"name"`
		HostName string   `json:"hostName"`
		Speakers []string `json:"speakers"`
	}{
		Name:     p.Name,
		HostName: p.HostName,
		Speakers: p.Speakers,
	})
}

// UnmarshalJSON decodes JSON to the Podcast struct
func (p *Podcast) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Name     string   `json:"name"`
		HostName string   `json:"hostName"`
		Speakers []string `json:"speakers"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	p.Name = tmp.Name
	p.HostName = tmp.HostName
	p.Speakers = tmp.Speakers
	return nil
}

// Function to take input from the user
func takeInput() Podcast {
	var podcast Podcast

	fmt.Print("Enter podcast name: ")
	fmt.Scanln(&podcast.Name)

	fmt.Print("Enter host name: ")
	fmt.Scanln(&podcast.HostName)

	fmt.Print("Enter speaker names (comma-separated): ")
	var speakerInput string
	fmt.Scanln(&speakerInput)
	podcast.Speakers = strings.Split(speakerInput, ",")

	return podcast
}

func TestPodcastJSON(t *testing.T) {
	// Create a sample podcast
	podcast := Podcast{
		Name:     "Sample Podcast",
		HostName: "John Doe",
		Speakers: []string{"Speaker 1", "Speaker 2"},
	}

	// Test MarshalJSON
	jsonData, err := json.Marshal(podcast)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}
	expectedJSON := `{"name":"Sample Podcast","hostName":"John Doe","speakers":["Speaker 1","Speaker 2"]}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON: %s, Got: %s", expectedJSON, jsonData)
	}

	// Test UnmarshalJSON
	var newPodcast Podcast
	if err := json.Unmarshal([]byte(expectedJSON), &newPodcast); err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}
	if newPodcast.Name != podcast.Name || newPodcast.HostName != podcast.HostName || len(newPodcast.Speakers) != len(podcast.Speakers) {
		t.Errorf("UnmarshalJSON: Expected %+v, Got: %+v", podcast, newPodcast)
	}
	for i := range newPodcast.Speakers {
		if newPodcast.Speakers[i] != podcast.Speakers[i] {
			t.Errorf("UnmarshalJSON: Expected %+v, Got: %+v", podcast, newPodcast)
		}
	}
}

func main() {
	fmt.Println("Welcome to the Podcast Streaming System.")
	fmt.Println("This program allows you to input details about podcasts, including the podcast name, host name, and speaker names.")

	// Take input from the user
	podcast := takeInput()

	// Encode the podcast struct to JSON
	jsonData, err := json.Marshal(podcast)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON encoded podcast:", string(jsonData))

	// Decode JSON to podcast struct
	var decodedPodcast Podcast
	if err := json.Unmarshal(jsonData, &decodedPodcast); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded podcast:", decodedPodcast)
}
