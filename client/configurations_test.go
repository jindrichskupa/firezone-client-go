package client

import (
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	configuration, err := client.GetConfiguration()
	if err != nil {
		t.Fatalf("Error getting configuration: %s", err)
	}
	if configuration == nil {
		t.Fatalf("Expected configuration, got nil")
	}
}
