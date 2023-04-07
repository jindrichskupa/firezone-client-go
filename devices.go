package firezone

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllDevices - Returns all user's Device
func (c *Client) GetAllDevices() (*[]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/users", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiDevices := ApiDevices{}
	err = json.Unmarshal(body, &apiDevices)
	if err != nil {
		return nil, err
	}

	return &apiDevices.data, nil
}

// GetDevice - Returns a specifc Device
func (c *Client) GetDevice(deviceId string) (*Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/users/%s", c.HostURL, deviceId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiDevice := ApiDevice{}
	err = json.Unmarshal(body, &apiDevice)
	if err != nil {
		return nil, err
	}

	return &apiDevice.data, nil
}

// CreateDevice - Create new Device
func (c *Client) CreateDevice(device Device) (*Device, error) {
	rb, err := json.Marshal(device)
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

	apiDevice := ApiDevice{}
	err = json.Unmarshal(body, &apiDevice)
	if err != nil {
		return nil, err
	}

	return &apiDevice.data, nil
}

// UpdateDevice - Updates an Device
func (c *Client) UpdateDevice(deviceId string, device Device) (*Device, error) {
	rb, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v0/users/%s", c.HostURL, deviceId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiDevice := ApiDevice{}
	err = json.Unmarshal(body, &apiDevice)
	if err != nil {
		return nil, err
	}

	return &apiDevice.data, nil
}

// DeleteDevice - Deletes an Device
func (c *Client) DeleteDevice(deviceId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v0/users/%s", c.HostURL, deviceId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted Device" {
		return errors.New(string(body))
	}

	return nil
}
