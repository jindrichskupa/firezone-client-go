package client

import (
	"os"
	"testing"
)

func TestGetAllRules(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	_, err = client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.1",
		PortRange:   "81",
		PortType:    "tcp",
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}

	rules, err := client.GetAllRules()
	if err != nil {
		t.Fatalf("Error getting all rules: %s", err)
	}
	if len(*rules) == 0 {
		t.Fatalf("Expected at least one rule, got %d", len(*rules))
	}
}

func TestCreateGlobalRule(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	rule, err := client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.2",
		PortRange:   "82",
		PortType:    "tcp",
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}
	newRule, err := client.GetRule(rule.ID)
	if err != nil {
		t.Fatalf("Error getting new rule: %s", err)
	}
	if newRule.Action != rule.Action {
		t.Fatalf("Expected action %s, got %s", rule.Action, newRule.Action)
	}
	if newRule.Destination != rule.Destination {
		t.Fatalf("Expected destination %s, got %s", rule.Destination, newRule.Destination)
	}
	if newRule.PortRange != rule.PortRange {
		t.Fatalf("Expected port range %s, got %s", rule.PortRange, newRule.PortRange)
	}
	if newRule.PortType != rule.PortType {
		t.Fatalf("Expected port type %s, got %s", rule.PortType, newRule.PortType)
	}
}

func TestCreateUserRule(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}
	rule, err := client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.3",
		PortRange:   "83",
		PortType:    "tcp",
		UserId:      user.ID,
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}
	newRule, err := client.GetRule(rule.ID)
	if err != nil {
		t.Fatalf("Error getting new rule: %s", err)
	}
	if newRule.Action != rule.Action {
		t.Fatalf("Expected action %s, got %s", rule.Action, newRule.Action)
	}
	if newRule.Destination != rule.Destination {
		t.Fatalf("Expected destination %s, got %s", rule.Destination, newRule.Destination)
	}
	if newRule.PortRange != rule.PortRange {
		t.Fatalf("Expected port range %s, got %s", rule.PortRange, newRule.PortRange)
	}
	if newRule.PortType != rule.PortType {
		t.Fatalf("Expected port type %s, got %s", rule.PortType, newRule.PortType)
	}
}

func TestUpdateGlobalRule(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	rule, err := client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.4",
		PortRange:   "84",
		PortType:    "tcp",
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}
	rule.Action = "reject"
	newRule, err := client.UpdateRule(rule.ID, *rule)
	if err != nil {
		t.Logf("Error updating rule: %s", err)
		return
	}
	if newRule.Action != rule.Action {
		t.Logf("Expected action %s, got %s", rule.Action, newRule.Action)
	}
	if newRule.Destination != rule.Destination {
		t.Logf("Expected destination %s, got %s", rule.Destination, newRule.Destination)
	}
	if newRule.PortRange != rule.PortRange {
		t.Logf("Expected port range %s, got %s", rule.PortRange, newRule.PortRange)
	}
	if newRule.PortType != rule.PortType {
		t.Logf("Expected port type %s, got %s", rule.PortType, newRule.PortType)
	}
}

func TestUpdateUserRule(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	user, err := client.CreateUser(User{
		Email: GenerateRandomEmail(),
		Role:  "admin",
	})
	if err != nil {
		t.Fatalf("Error creating user: %s", err)
	}
	rule, err := client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.5",
		PortRange:   "85",
		PortType:    "tcp",
		UserId:      user.ID,
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}
	rule.Action = "reject"
	newRule, err := client.UpdateRule(rule.ID, *rule)
	if err != nil {
		t.Logf("Error updating rule: %s", err)
		return
	}
	if newRule.Action != rule.Action {
		t.Logf("Expected action %s, got %s", rule.Action, newRule.Action)
	}
	if newRule.Destination != rule.Destination {
		t.Logf("Expected destination %s, got %s", rule.Destination, newRule.Destination)
	}
	if newRule.PortRange != rule.PortRange {
		t.Logf("Expected port range %s, got %s", rule.PortRange, newRule.PortRange)
	}
	if newRule.PortType != rule.PortType {
		t.Logf("Expected port type %s, got %s", rule.PortType, newRule.PortType)
	}
}

func TestDeleteRule(t *testing.T) {
	client, err := NewClient(os.Getenv("FIREZONE_ENDPOINT"), os.Getenv("FIREZONE_API_KEY"))

	rule, err := client.CreateRule(Rule{
		Action:      "accept",
		Destination: "127.0.0.6",
		PortRange:   "86",
		PortType:    "tcp",
	})
	if err != nil {
		t.Fatalf("Error creating rule: %s", err)
	}
	err = client.DeleteRule(rule.ID)
	if err != nil {
		t.Fatalf("Error deleting rule: %s", err)
	}
	_, err = client.GetRule(rule.ID)
	if err == nil {
		t.Fatalf("Expected error getting rule, got nil")
	}
}
