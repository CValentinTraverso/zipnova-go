package zipnova

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	CountryArgentina = "zipnova.com.ar"
	CountryChile     = "zipnova.cl"
	CountryMexico    = "zipnova.com.mx"

	defaultVersion = "v2"
)

type Country string

type Client struct {
	httpClient  *http.Client
	baseURL     string
	apiToken    string
	apiSecret   string
}

func NewClient(apiToken, apiSecret string, country Country) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    fmt.Sprintf("https://api.%s/%s", country, defaultVersion),
		apiToken:   apiToken,
		apiSecret:  apiSecret,
	}
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiToken, c.apiSecret)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}
