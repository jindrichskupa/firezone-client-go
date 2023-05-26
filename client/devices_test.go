package client

import (
	"os"
	"testing"
)

func TestGetAllDevices(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})

	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	_, err = client.CreateDevice(Device{
		Description: "Test Device",
		Name:        "test-device",
		PublicKey:   GenerateRandomString(32),
		UserId:      user.ID,
	})

	if err != nil {
		t.Fatalf("Error creating device: %s", err)
	}

	devices, err := client.GetAllDevices()
	if err != nil {
		t.Fatalf("Error getting all devices: %s", err)
	}
	if len(*devices) == 0 {
		t.Fatalf("Expected at least one device, got %d", len(*devices))
	}
}
