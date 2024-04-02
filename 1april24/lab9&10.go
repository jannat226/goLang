package main

import (
	"fmt"
	"time"
)

// Podcast represents a podcast episode
type Podcast struct {
	Name     string
	HostName string
	Speakers []string
}

// Streamer interface to represent entities that can stream content
type Streamer interface {
	Stream()
}

// Implementing the Streamer interface for Podcast
func (p Podcast) Stream() {
	fmt.Printf("Streaming podcast %s hosted by %s, with speakers %s\n", p.Name, p.HostName, p.Speakers)
	time.Sleep(5 * time.Second) // Simulate streaming time
}

func main() {
	fmt.Println("Welcome to the Podcast Streaming System.")

	// Create a channel to handle streaming
	streamChannel := make(chan Streamer)

	// Start streaming podcasts concurrently
	go func() {
		for {
			select {
			case streamer := <-streamChannel:
				streamer.Stream()
			}
		}
	}()

	// Simulate streaming of multiple podcasts
	for i := 1; i <= 5; i++ {
		podcast := Podcast{
			Name:     fmt.Sprintf("Podcast %d", i),
			HostName: fmt.Sprintf("Host %d", i),
			Speakers: []string{fmt.Sprintf("Speaker %d", i)},
		}
		streamChannel <- podcast
	}

	// Wait for a few seconds before exiting to allow streaming to finish
	time.Sleep(10 * time.Second)
}
