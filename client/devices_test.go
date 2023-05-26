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

func TestCreateDevice(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})

	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	device, err := client.CreateDevice(Device{
		Description: "Test Device",
		Name:        "test-device",
		PublicKey:   GenerateRandomString(32),
		UserId:      user.ID,
	})

	if err != nil {
		t.Fatalf("Error creating device: %s", err)
	}

	newDevice, err := client.GetDevice(device.ID)
	if err != nil {
		t.Fatalf("Error getting new device: %s", err)
	}
	if newDevice.Name != device.Name {
		t.Fatalf("Expected name %s, got %s", device.Name, newDevice.Name)
	}
}

func TestDeleteDevice(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})

	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	device, err := client.CreateDevice(Device{
		Description: "Test Device",
		Name:        "test-device",
		PublicKey:   GenerateRandomString(32),
		UserId:      user.ID,
	})

	if err != nil {
		t.Fatalf("Error creating device: %s", err)
	}

	err = client.DeleteDevice(device.ID)
	if err != nil {
		t.Fatalf("Error deleting device: %s", err)
	}
}

func TestUpdateDevice(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})

	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}

	device, err := client.CreateDevice(Device{
		Description: "Test Device",
		Name:        "test-device",
		PublicKey:   GenerateRandomString(32),
		UserId:      user.ID,
	})

	if err != nil {
		t.Fatalf("Error creating device: %s", err)
	}

	device.Description = "New Description"
	newDevice, err := client.UpdateDevice(device.ID, *device)
	if err != nil {
		t.Fatalf("Error updating device: %s", err)
	}
	if newDevice.Description != device.Description {
		t.Fatalf("Expected description %s, got %s", device.Description, newDevice.Description)
	}
}