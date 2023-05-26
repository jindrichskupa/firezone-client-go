package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	fz "github.com/jindrichskupa/firezone-client-go/client"
)

func publicKey() string {
	// Generate a random byte slice of length 33
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println("error generating random bytes:", err)
		return ""
	}

	// Encode the byte slice using base64
	encoded := base64.StdEncoding.EncodeToString(bytes)

	return encoded
}

func main() {
	fmt.Println("Firezone API client example")
	fmt.Println("FIREZONE_ENDPOINT: ", os.Getenv("FIREZONE_ENDPOINT"))

	client, err := fz.NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	if err != nil {
		fmt.Println(err)
	}

	// Get all users
	fmt.Println("Users: ")
	users, err := client.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fz.PrintUsers(users))

	client.DeleteUser("test5@example.com")
	user, err := client.CreateUser(fz.User{
		Email: "test5@example.com",
		Role:  "admin",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.PrintUser())
	}

	user.Role = "unprivileged"

	user, err = client.UpdateUser(user.ID, *user)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.PrintUser())
	}

	user, err = client.GetUser(user.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.PrintUser())
	}

	user, err = client.GetUser(user.Email)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.PrintUser())
	}

	// Get all devices

	devices, err := client.GetAllDevices()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Devices: ")
	fmt.Println(fz.PrintDevices(devices))

	device, err := client.CreateDevice(fz.Device{
		Name:        "test5",
		Description: "test5",
		UserId:      user.ID,
		PublicKey:   publicKey(),
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(device.PrintDevice())
	}

	device, err = client.GetDevice(device.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(device.PrintDevice())
	}

	client.DeleteDevice(device.ID)
	client.DeleteUser("test5@example.com")

	// Get all rules

	rules, err := client.GetAllRules()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Rules: ")
	fmt.Println(fz.PrintRules(rules))

	rule, err := client.CreateRule(fz.Rule{
		Action:      "drop",
		Destination: "1.1.1.0/24",
		PortRange:   "80",
		PortType:    "tcp",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rule.PrintRule())
	}
	client.DeleteRule(rule.ID)

	// Get configuration

	configuration, err := client.GetConfiguration()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Configuration: ")
	fmt.Println(configuration.PrintConfiguration())
}
