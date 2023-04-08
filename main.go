package main

import (
	"fmt"
	"os"

	fz "github.com/jindrichskupa/firezone-client-go/client"
)

func main() {
	client, err := fz.NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	if err != nil {
		fmt.Println(err)
	}

	// Get all users

	users, err := client.GetAllUsers()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Users: ")
	fmt.Println(fz.PrintUsers(users))

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

	client.DeleteUser("test5@example.com")

	// Get all devices

	devices, err := client.GetAllDevices()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Devices: ")
	fmt.Println(fz.PrintDevices(devices))

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

	// Get configuration

	configuration, err := client.GetConfiguration()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Configuration: ")
	fmt.Println(configuration.PrintConfiguration())
}
