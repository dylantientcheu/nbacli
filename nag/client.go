package nag

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	// DefaultBaseURL sets default base URL for request to NBA stats.
	DefaultBaseURL = &url.URL{
		Host:   "stats.nba.com",
		Scheme: "https",
		Path:   "/stats",
	}
	// DefaultStatsHeader sets default headers for request to NBA stats.
	// no idea which is necessary and which is not
	DefaultStatsHeader = http.Header{
		"Host":               []string{"stats.nba.com"},
		"Referer":            []string{"https://stats.nba.com"},
		"User-Agent":         []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0"},
		"Connection":         []string{"keep-alive"},
		"Pragma":             []string{"no-cache"},
		"Cache-Control":      []string{"no-cache"},
		"Accept":             []string{"application/json", "text/plain", "*/*"},
		"Accept-Encoding":    []string{"gzip", "deflate", "br"},
		"Accept-Language":    []string{"en-US,en;q=0.9"},
		"x-nba-stats-origin": []string{"stats"},
		"x-nba-stats-token":  []string{"true"},
	}
)

// Client contains the base URL to send request to and the HTTP client being used.
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// NewDefaultClient uses stdlib default HTTP client to make request to
// default NBA stats endpoint.
func NewDefaultClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: http.DefaultClient,
	}
}

// Do sends request to NBA stats endpoint and unpacks the received response.
func (c *Client) Do(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s status code: %d", req.URL.String(), res.StatusCode)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	gr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, nil
	}
	defer gr.Close()

	b, err = io.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return b, nil
}
