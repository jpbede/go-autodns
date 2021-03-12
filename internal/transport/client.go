package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Client is the http transport client for netpalm. It handles the authentication
type Client struct {
	BaseURL     string
	Credentials *APICredentials

	httpClient *http.Client
	logger     *zerolog.Logger
}

// NewClient returns a new Transport HTTP client
func NewClient(baseURL string, hc *http.Client) *Client {
	if hc == nil {
		hc = &http.Client{}
	}

	c := Client{
		BaseURL:    baseURL,
		httpClient: hc,
	}
	return &c
}

// Get executes a GET request
func (c *Client) Get(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodGet, path, out, opts...)
}

// Post executes a POST request
func (c *Client) Post(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPost, path, out, opts...)
}

// Put executes a PUT request
func (c *Client) Put(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPut, path, out, opts...)
}

// Patch executes a PATCH request
func (c *Client) Patch(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodPatch, path, out, opts...)
}

// Delete executes a DELETE request
func (c *Client) Delete(ctx context.Context, path string, out interface{}, opts ...RequestOption) error {
	return c.doRequest(ctx, http.MethodDelete, path, out, opts...)
}

// doRequest does the actual request
func (c *Client) doRequest(ctx context.Context, method string, path string, out interface{}, opts ...RequestOption) error {
	// create a new request
	path = strings.TrimPrefix(path, "/")
	req, err := http.NewRequest(method, c.BaseURL+"/"+path, nil)
	if err != nil {
		return err
	}

	// add client logger if zerolog logger is there
	if c.logger != nil {
		c.httpClient.Transport = &ClientLogging{Logger: c.logger}
	}

	// add auth headers
	req.SetBasicAuth(c.Credentials.Username, c.Credentials.Password)
	req.Header.Add("X-Domainrobot-Context", strconv.Itoa(c.Credentials.Context))

	req.Header.Set("User-Agent", "go-autodns/0.0.0")

	// run options
	for i := range opts {
		if err := opts[i](req); err != nil {
			return err
		}
	}
	req = req.WithContext(ctx)

	// run request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		bytes, _ := ioutil.ReadAll(res.Body)
		return errors.New(string(bytes))
	}

	// marshal response to given interface
	if out != nil {
		if w, ok := out.(io.Writer); ok {
			_, err := io.Copy(w, res.Body)
			return err
		}

		if err := json.NewDecoder(res.Body).Decode(out); err != nil {
			return err
		}
	}

	return nil
}

// SetLogger sets a zerolog logger for request and response logging
func (c *Client) SetLogger(logger *zerolog.Logger) {
	c.logger = logger
}

// HasCredentials checks if the transport client has credentials
func (c *Client) HasCredentials() bool {
	return c.Credentials != nil
}
