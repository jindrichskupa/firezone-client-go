package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetConfiguration -
func (c *Client) GetConfiguration() (*Configuration, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/configuration", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiConfiguration := ApiConfiguration{}
	err = json.Unmarshal(body, &apiConfiguration)
	if err != nil {
		return nil, err
	}

	return &apiConfiguration.Data, nil
}

// UpdateConfiguration  -
func (c *Client) UpdateConfiguration(configuration Configuration) (*Configuration, error) {
	rb, err := json.Marshal(configuration)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v0/configuration", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiConfiguration := ApiConfiguration{}
	err = json.Unmarshal(body, &apiConfiguration)
	if err != nil {
		return nil, err
	}

	return &apiConfiguration.Data, nil
}
