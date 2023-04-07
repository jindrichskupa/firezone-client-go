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

	users, err := client.GetAllUsers()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}
