package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetConfiguration -
func (c *Client) GetConfiguration() (*Configuration, error) {
	url := fmt.Sprintf("%s/v0/configuration", c.HostURL)
	log.Printf("GetConfiguration %s", url)

	req, err := http.NewRequest("GET", url, nil)
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

	url := fmt.Sprintf("%s/v0/configuration", c.HostURL)
	log.Printf("UpdateConfiguration %s", url)

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(rb)))
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
