// Package api provides the HTTP client for communicating with the mkt36z Workers API.
package api

import (
	"context"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/mkt36z/cli/internal/version"
)

// Client is the HTTP client for the mkt36z Workers API.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client

	// MaxRetries is the number of retries for transient failures (default 3).
	MaxRetries int
}

// NewClient creates a new API client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		MaxRetries: 3,
	}
}

// NewStreamingClient creates a client configured for long-running SSE streams.
func NewStreamingClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		MaxRetries: 3,
	}
}

// Do executes an HTTP request with retries, auth headers, and request ID tracking.
func (c *Client) Do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	var lastErr error

	for attempt := range c.MaxRetries + 1 {
		req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}

		reqID := newRequestID()
		c.setHeaders(req, reqID)

		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = &NetworkError{Err: err, RequestID: reqID}
			if attempt < c.MaxRetries {
				backoff(ctx, attempt)
				continue
			}
			return nil, lastErr
		}

		// Retry on 5xx (server errors) and 429 (rate limit)
		if resp.StatusCode >= 500 || resp.StatusCode == http.StatusTooManyRequests {
			resp.Body.Close()
			lastErr = classifyError(resp.StatusCode, reqID)
			if attempt < c.MaxRetries {
				backoff(ctx, attempt)
				continue
			}
			return nil, lastErr
		}

		// Check for usage warning header (print to stderr)
		if warning := resp.Header.Get("X-Usage-Warning"); warning != "" {
			fmt.Fprintf(os.Stderr, "Warning: %s. Upgrade: mkt36z usage upgrade\n", warning)
		}

		// Client errors (4xx) are not retried
		if resp.StatusCode >= 400 {
			resp.Body.Close()
			return nil, classifyError(resp.StatusCode, reqID)
		}

		return resp, nil
	}

	return nil, lastErr
}

func (c *Client) setHeaders(req *http.Request, reqID string) {
	info := version.Get()
	req.Header.Set("User-Agent", fmt.Sprintf("mkt36z-cli/%s", info.Version))
	req.Header.Set("X-Request-ID", reqID)
	req.Header.Set("X-CLI-Version", info.Version)
	req.Header.Set("Content-Type", "application/json")

	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
}

// newRequestID generates a request ID with a timestamp prefix.
func newRequestID() string {
	now := time.Now().UnixMilli()
	r := rand.Uint32()
	return fmt.Sprintf("req_%x%08x", now, r)
}

// backoff waits with exponential backoff + jitter. Respects context cancellation.
func backoff(ctx context.Context, attempt int) {
	base := time.Duration(1<<uint(attempt)) * time.Second // 1s, 2s, 4s
	jitter := time.Duration(rand.Int64N(int64(base / 2)))
	wait := base + jitter

	select {
	case <-ctx.Done():
	case <-time.After(wait):
	}
}
