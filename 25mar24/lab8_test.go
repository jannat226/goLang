package main
import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestPodcastJSONMarshalling(t *testing.T) {
	// Create a Podcast struct with sample data
	podcast := Podcast{
		Title:       "Sample Podcast",
		Host:        "John Doe",
		Episode:     10,
		Description: "This is a sample podcast episode.",
	}

	// Marshal Podcast struct into JSON
	jsonData, err := json.Marshal(podcast)
	if err != nil {
		t.Errorf("Error marshalling JSON: %v", err)
	}

	// Expected JSON
	expectedJSON := `{"title":"Sample Podcast","host":"John Doe","episode":10,"description":"This is a sample podcast episode."}`

	// Compare the marshalled JSON with expected JSON
	if !bytes.Equal(jsonData, []byte(expectedJSON)) {
		t.Errorf("Unexpected JSON result:\nExpected: %s\nGot: %s", expectedJSON, jsonData)
	}

	// Unmarshal JSON back into a Podcast struct
	var newPodcast Podcast
	err = json.Unmarshal(jsonData, &newPodcast)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %v", err)
	}

	// Compare the original struct with the unmarshalled struct
	if newPodcast != podcast {
		t.Errorf("Unmarshalled struct is different from original:\nOriginal: %+v\nUnmarshalled: %+v", podcast, newPodcast)
	}
}
