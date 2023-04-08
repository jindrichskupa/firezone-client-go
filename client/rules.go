package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetAllRules - Returns all user's Rule
func (c *Client) GetAllRules() (*[]Rule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/rules", c.HostURL), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId), nil)
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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v0/rules", c.HostURL), strings.NewReader(string(rb)))
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

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId), strings.NewReader(string(rb)))
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v0/rules/%s", c.HostURL, ruleId), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted Rule" {
		return errors.New(string(body))
	}

	return nil
}
