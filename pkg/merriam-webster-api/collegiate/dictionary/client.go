package dictionary

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// A Client is an HTTP based client lookup word definitions
// using Merriam-Websters Collegiate Dictionary API.
type Client struct {
	baseURL    *url.URL
	userAgent  string
	httpClient *http.Client

	// requests made
	// successful requests made
}

// NewClient returns a new Client given an API key, a user agent, and
// an HTTP client.
func NewClient(apiKey string, userAgent string, client *http.Client) (*Client, error) {
	baseUrl, err := url.Parse("https://dictionaryapi.com/api/v3/references/collegiate/json")
	if err != nil {
		return nil, err
	}

	q := baseUrl.Query()
	q.Set("key", apiKey)
	baseUrl.RawQuery = q.Encode()

	return &Client{
		baseURL:    baseUrl,
		userAgent:  userAgent,
		httpClient: client,
	}, nil
}

// SetBaseURL sets the base URL used by the client
func (c *Client) SetBaseURL(baseURL *url.URL) {
	c.baseURL = baseURL
}

// Lookup returns the dictionary response if successful. On failure the
// reason is returned.
func (c *Client) Lookup(ctx context.Context, word string) (*Response, error) {
	req, err := c.newRequest("GET", word, nil)
	if err != nil {
		return nil, err
	}

	var resp *Response
	_, err = c.do(ctx, req, resp)
	return resp, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
