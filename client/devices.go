package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllDevices - Returns all user's Device
func (c *Client) GetAllDevices() (*[]Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/devices", c.HostURL), nil)
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

	return &apiDevices.Data, nil
}

// GetDevice - Returns a specifc Device
func (c *Client) GetDevice(deviceId string) (*Device, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId), nil)
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

	return &apiDevice.Data, nil
}

// CreateDevice - Create new Device
func (c *Client) CreateDevice(device Device) (*Device, error) {
	rb, err := json.Marshal(CreateDevice{Device: struct {
		AllowedIPs             []string `json:"allowed_ips"`
		Description            string   `json:"description"`
		DNS                    []string `json:"dns"`
		Endpoint               string   `json:"endpoint"`
		IPv4                   string   `json:"ipv4"`
		IPv6                   string   `json:"ipv6"`
		MTU                    int      `json:"mtu"`
		Name                   string   `json:"name"`
		PersistentKeepalive    int      `json:"persistent_keepalive"`
		PresharedKey           string   `json:"preshared_key"`
		PublicKey              string   `json:"public_key"`
		UseDefaultAllowedIPs   bool     `json:"use_default_allowed_ips"`
		UseDefaultDNS          bool     `json:"use_default_dns"`
		UseDefaultEndpoint     bool     `json:"use_default_endpoint"`
		UseDefaultMTU          bool     `json:"use_default_mtu"`
		UseDefaultPersistentKA bool     `json:"use_default_persistent_keepalive"`
		UserId                 string   `json:"user_id"`
	}{
		AllowedIPs:          device.AllowedIPs,
		Description:         device.Description,
		DNS:                 device.DNS,
		Endpoint:            device.Endpoint,
		IPv4:                device.IPv4,
		IPv6:                device.IPv6,
		MTU:                 device.MTU,
		Name:                device.Name,
		PersistentKeepalive: device.PersistentKeepalive,
		UserId:              device.UserId,
	}})

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v0/devices", c.HostURL), strings.NewReader(string(rb)))
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

	return &apiDevice.Data, nil
}

// UpdateDevice - Updates an Device
func (c *Client) UpdateDevice(deviceId string, device Device) (*Device, error) {
	rb, err := json.Marshal(device)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId), strings.NewReader(string(rb)))
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

	return &apiDevice.Data, nil
}

// DeleteDevice - Deletes an Device
func (c *Client) DeleteDevice(deviceId string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId), nil)
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
