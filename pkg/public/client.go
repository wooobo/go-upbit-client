package public

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	BaseUrl string
	Version string
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(client Config) *Client {
	return &Client{
		baseURL: fmt.Sprintf("%s%s", client.BaseUrl, client.Version),
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) Get(ctx context.Context, path string, params url.Values, v interface{}) error {
	return c.do(ctx, http.MethodGet, path, params, nil, v)
}

func (c *Client) do(ctx context.Context, method, path string, params url.Values, body interface{}, v interface{}) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			err = fmt.Errorf("closing response body: %w", err)
		}
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("decoding response body: %w", err)
		}
	}

	return nil
}

func (c *Client) toUrl(u string) string {
	return fmt.Sprintf("%s%s", c.baseURL, u)
}
