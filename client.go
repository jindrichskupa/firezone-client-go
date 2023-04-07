package firezone

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Firezone URL
const HostURL string = "http://localhost:19090"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	ApiKey     string
}

// NewClient -
func NewClient(host, apiKey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Firezone URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If apiKey not provided, return empty client
	if apiKey == nil {
		return &c, nil
	}

	c.ApiKey = *apiKey

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	token := c.ApiKey

	req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
