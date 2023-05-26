package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetAllRules - Returns all user's Rule
func (c *Client) GetAllRules() (*[]Rule, error) {
	url := fmt.Sprintf("%s/v0/rules", c.HostURL)
	log.Printf("GetAllRules %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiRules := ApiRules{}
	err = json.Unmarshal(body, &apiRules)
	if err != nil {
		return nil, err
	}

	return &apiRules.Data, nil
}

// GetRule - Returns a specifc Rule
func (c *Client) GetRule(ruleId string) (*Rule, error) {
	url := fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId)
	log.Printf("GetRule %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiRule := ApiRule{}
	err = json.Unmarshal(body, &apiRule)
	if err != nil {
		return nil, err
	}

	return &apiRule.Data, nil
}

// CreateRule - Create new Rule
func (c *Client) CreateRule(rule Rule) (*Rule, error) {
	rb, err := json.Marshal(CreateRule{Rule: struct {
		Action      string `json:"action"`
		Destination string `json:"destination"`
		PortRange   string `json:"port_range"`
		PortType    string `json:"port_type"`
		UserId      string `json:"user_id,omitempty"`
	}{
		Action:      rule.Action,
		Destination: rule.Destination,
		PortRange:   rule.PortRange,
		PortType:    rule.PortType,
		UserId:      rule.UserId,
	}})

	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/rules", c.HostURL)
	log.Printf("CreateRule %s", url)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiRule := ApiRule{}
	err = json.Unmarshal(body, &apiRule)
	if err != nil {
		return nil, err
	}

	return &apiRule.Data, nil
}

// UpdateRule - Updates an Rule
func (c *Client) UpdateRule(ruleId string, rule Rule) (*Rule, error) {
	rb, err := json.Marshal(rule)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId)
	log.Printf("UpdateRule %s", url)

	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiRule := ApiRule{}
	err = json.Unmarshal(body, &apiRule)
	if err != nil {
		return nil, err
	}

	return &apiRule.Data, nil
}

// DeleteRule - Deletes an Rule
func (c *Client) DeleteRule(ruleId string) error {
	url := fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId)
	log.Printf("DeleteRule %s", url)

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
