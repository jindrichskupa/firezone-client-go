package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetAllDevices - Returns all user's Device
func (c *Client) GetAllDevices() (*[]Device, error) {
	url := fmt.Sprintf("%s/v0/devices", c.HostURL)
	log.Printf("GetAllDevices %s", url)

	req, err := http.NewRequest("GET", url, nil)
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
	url := fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId)
	log.Printf("GetDevice %s", url)

	req, err := http.NewRequest("GET", url, nil)
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
		AllowedIPs             []string `json:"allowed_ips,omitempty"`
		Description            string   `json:"description"`
		DNS                    []string `json:"dns,omitempty"`
		Endpoint               string   `json:"endpoint,omitempty"`
		IPv4                   string   `json:"ipv4,omitempty"`
		IPv6                   string   `json:"ipv6,omitempty"`
		MTU                    int      `json:"mtu,omitempty"`
		Name                   string   `json:"name"`
		PersistentKeepalive    int      `json:"persistent_keepalive,omitempty"`
		PresharedKey           string   `json:"preshared_key,omitempty"`
		PublicKey              string   `json:"public_key"`
		UseDefaultAllowedIPs   bool     `json:"use_default_allowed_ips,omitempty"`
		UseDefaultDNS          bool     `json:"use_default_dns,omitempty"`
		UseDefaultEndpoint     bool     `json:"use_default_endpoint,omitempty"`
		UseDefaultMTU          bool     `json:"use_default_mtu,omitempty"`
		UseDefaultPersistentKA bool     `json:"use_default_persistent_keepalive,omitempty"`
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
		PublicKey:           device.PublicKey,
		PersistentKeepalive: device.PersistentKeepalive,
		UserId:              device.UserId,
	}})

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/devices", c.HostURL)
	log.Printf("CreateDevice %s", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
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

	url := fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId)
	log.Printf("UpdateDevice %s", url)

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(rb)))
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
	url := fmt.Sprintf("%s/v0/devices/%s", c.HostURL, deviceId)
	log.Printf("DeleteDevice %s", url)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
