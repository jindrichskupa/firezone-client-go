package client

import (
	"os"
	"testing"
)

func TestCreateUser(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}
	newUser, err := client.GetUser(user.ID)
	if err != nil {
		t.Fatalf("Error getting new user: %s", err)
	}
	if newUser.Email != user.Email {
		t.Fatalf("Expected email %s, got %s", user.Email, newUser.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}
	user.Role = "unprivileged"
	newUser, err := client.UpdateUser(user.ID, *user)
	if err != nil {
		t.Fatalf("Error updating user: %s", err)
	}
	if newUser.Role != user.Role {
		t.Fatalf("Expected role %s, got %s", user.Role, newUser.Role)
	}
}

func TestGetAllUsers(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	users, err := client.GetAllUsers()
	if err != nil {
		t.Fatalf("Error getting all users: %s", err)
	}
	if len(*users) == 0 {
		t.Fatalf("Expected at least one user, got %d", len(*users))
	}
}

func TestDeleteUser(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}
	err = client.DeleteUser(user.ID)
	if err != nil {
		t.Fatalf("Error deleting user: %s", err)
	}
	_, err = client.GetUser(user.ID)
	if err == nil {
		t.Fatalf("Expected error getting user, got none")
	}
}
