package private

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	PublicApiKey string
	SecretApiKey string
	BaseUrl      string
	Version      string
}

type Client struct {
	baseURL    string
	jwtManager *JWTManager
	httpClient *http.Client
}

func NewClient(client Config) *Client {
	return &Client{
		baseURL:    client.BaseUrl + client.Version,
		jwtManager: NewJWT(client.PublicApiKey, client.SecretApiKey),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) Get(ctx context.Context, path string, body url.Values, v interface{}) error {
	return c.do(ctx, http.MethodGet, path, body, nil, v)
}

func (c *Client) Post(ctx context.Context, path string, body url.Values, v interface{}) error {
	return c.do(ctx, http.MethodPost, path, nil, body, v)
}

func (c *Client) Delete(ctx context.Context, path string, body url.Values, v interface{}) error {
	return c.do(ctx, http.MethodDelete, path, nil, body, v)
}

func (c *Client) do(ctx context.Context, method, path string, params url.Values, body url.Values, v interface{}) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		bodyReader = strings.NewReader(body.Encode())
	}

	urlPath := u.String()
	req, err := http.NewRequestWithContext(ctx, method, urlPath, bodyReader)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	if method == http.MethodPost || method == http.MethodDelete {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", c.jwtManager.CreateTokenWithQuery(body))
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("Authorization", c.jwtManager.CreateTokenWithQuery(params))
	}

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
