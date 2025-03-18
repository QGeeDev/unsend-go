package unsend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Client struct {
	Client   *http.Client
	ApiKey   string
	BaseUrl  *url.URL
	Contacts Contacts
	Domains  Domains
	Emails   Emails
}

func NewClient() (*Client, error) {
	ApiKey := os.Getenv(ENV_KEY_API_KEY)
	baseURL := GetEnvOrDefault(ENV_KEY_BASE_URL, DEFAULT_BASE_URL)

	if strings.TrimSpace(ApiKey) == "" {
		return nil, fmt.Errorf("no value found for API Key")
	}

	baseUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		ApiKey: ApiKey,
		Client: &http.Client{
			Timeout: time.Second * 30,
			Transport: &UnsendTransport{
				ApiKey: ApiKey,
			},
		},
		BaseUrl: baseUrl,
	}

	client.Contacts = &ContactsImpl{Client: client}
	client.Domains = &DomainsImpl{Client: client}
	client.Emails = &EmailsImpl{Client: client}

	return client, nil
}

func (c *Client) NewRequest(method, urlAsString string, body interface{}) (*http.Request, error) {
	url, err := c.BaseUrl.Parse(urlAsString)
	if err != nil {
		return nil, err
	}

	var requestBody io.Reader

	if body != nil {
		serialized, _ := json.Marshal(body)
		requestBody = bytes.NewReader(serialized)
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(context.Background(), method, url.String(), requestBody)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Execute(req *http.Request, result interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("received non-2xx response: %d - %s", resp.StatusCode, string(respBody))
	}

	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func GetEnvOrDefault(envVariable string, defaultValue string) string {
	val, ok := os.LookupEnv(envVariable)
	if !ok {
		fmt.Printf("%s not set. Using default value of %s\n", envVariable, defaultValue)
		return defaultValue
	} else {
		return val
	}
}
