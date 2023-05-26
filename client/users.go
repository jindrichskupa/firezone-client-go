package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetAllUsers - Returns all user's User
func (c *Client) GetAllUsers() (*[]User, error) {
	url := fmt.Sprintf("%s/v0/users", c.HostURL)
	log.Printf("GetAllUsers %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiUsers := ApiUsers{}
	err = json.Unmarshal(body, &apiUsers)
	if err != nil {
		return nil, err
	}

	return &apiUsers.Data, nil
}

// GetUser - Returns a specifc User
func (c *Client) GetUser(userId string) (*User, error) {
	url := fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId)
	log.Printf("GetUser %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiUser := ApiUser{}
	err = json.Unmarshal(body, &apiUser)
	if err != nil {
		return nil, err
	}

	return &apiUser.Data, nil
}

// CreateUser - Create new User
func (c *Client) CreateUser(user User) (*User, error) {
	rb, err := json.Marshal(CreateUser{User: struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		Email: user.Email,
		Role:  user.Role,
	}})

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/users", c.HostURL)
	log.Printf("CreateUser %s", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiUser := ApiUser{}
	err = json.Unmarshal(body, &apiUser)
	if err != nil {
		return nil, err
	}
	return &apiUser.Data, nil
}

// UpdateUser - Updates an User
func (c *Client) UpdateUser(userId string, user User) (*User, error) {
	rb, err := json.Marshal(CreateUser{User: struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		Email: user.Email,
		Role:  user.Role,
	}})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId)
	log.Printf("UpdateUser %s", url)

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiUser := ApiUser{}
	err = json.Unmarshal(body, &apiUser)
	if err != nil {
		return nil, err
	}

	return &apiUser.Data, nil
}

// DeleteUser - Deletes an User
func (c *Client) DeleteUser(userId string) error {
	url := fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId)
	log.Printf("DeleteUser %s", url)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
