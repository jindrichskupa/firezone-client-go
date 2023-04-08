package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllUsers - Returns all user's User
func (c *Client) GetAllUsers() (*[]User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/users", c.HostURL), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId), nil)
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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v0/users", c.HostURL), strings.NewReader(string(rb)))
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

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId), strings.NewReader(string(rb)))
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v0/users/%s", c.HostURL, userId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted User" {
		return errors.New(string(body))
	}

	return nil
}
