package macaddress

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"

	// "fmt"
	"io"
	"io/ioutil"

	// "log"
	"net/http"
	"net/url"
	// "strings"
)

// Client for interacting with macaddress API.
type Client struct {
	BaseURL    *url.URL
	APIKey     string
	httpClient *http.Client
	Output     string
}

// NewClient is constructor
func NewClient(key string) *Client {
	u, _ := url.Parse("https://api.macaddress.io/v1")
	return &Client{
		BaseURL:    u,
		httpClient: &http.Client{},
		APIKey:     key,
		Output:     "json",
	}
}

// sendRequest sends a HTTP request to the macaddress API.
func (c *Client) sendRequest(method, term string, body io.Reader) (*http.Response, error) {
	// Compose URL
	rel := &url.URL{}
	targetURL := c.BaseURL.ResolveReference(rel)

	// Write body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// New HTTP GET request

	req, err := http.NewRequest(method, targetURL.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// Add api key to query
	if c.APIKey != "" {
		q := url.Values{}
		q.Add("apiKey", c.APIKey)
		q.Add("output", c.Output)
		q.Add("search", term)
		req.URL.RawQuery = q.Encode()
	}

	// log.Printf("Doing request: %s", targetURL.String())
	return (c.httpClient).Do(req)
}

// Search term: MAC address or OUI
func (c *Client) Search(term string) (*Response, error) {
	httpResp, err := c.sendRequest("GET", term, nil)
	if err != nil {
		return nil, err
	}
	result := &Response{}

	if httpResp.StatusCode == http.StatusOK {
		defer httpResp.Body.Close()
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	var ErrNotFound = errors.New("Error with status code " + strconv.Itoa(httpResp.StatusCode))
	return nil, ErrNotFound
}

// GetVendor vendor company name only, in text format. return "" for any errors
func (c *Client) GetVendor(term string) string {
	c.Output = "vendor"
	result := ""
	httpResp, err := c.sendRequest("GET", term, nil)
	if err != nil {
		return result
	}
	if httpResp.StatusCode == http.StatusOK {
		defer httpResp.Body.Close()
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return result
		}
		result = string(body)
	}
	return result
}
