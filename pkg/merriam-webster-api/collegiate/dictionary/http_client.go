package dictionary

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// A HTTPClient is an HTTP based client lookup word definitions
// using Merriam-Websters Collegiate Dictionary API.
type HTTPClient struct {
	BaseURL           *url.URL
	apiKey, userAgent string
	httpClient        *http.Client
}

// NewClient returns a new HTTPClient given an API key, a user agent, and
// an HTTP client.
func NewClient(apiKey string, userAgent string, client *http.Client) (*HTTPClient, error) {
	baseUrl, _ := url.Parse("https://dictionaryapi.com/api/v3/references/collegiate/json")

	return &HTTPClient{
		BaseURL:    baseUrl,
		apiKey:     apiKey,
		userAgent:  userAgent,
		httpClient: client,
	}, nil
}

// Lookup returns the dictionary response if successful. On failure the
// reason is returned.
func (c *HTTPClient) Lookup(ctx context.Context, word string) ([]*Response, error) {
	req, err := c.newRequest("GET", c.BaseURL.Path+"/"+word, nil)
	if err != nil {
		return nil, err
	}

	resp := make([]*Response, 0)
	_, err = c.do(ctx, req, &resp)
	return resp, err
}

func (c *HTTPClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	q := u.Query()
	q.Set("key", c.apiKey)
	u.RawQuery = q.Encode()

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

func (c *HTTPClient) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
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
