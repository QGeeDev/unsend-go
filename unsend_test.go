package unsend_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/QGeeDev/unsend-go"
)

func TestNewClient(t *testing.T) {
	baseURL := "https://api.example.com"

	os.Setenv(unsend.ENV_KEY_API_KEY, "test-api-key")
	defer os.Unsetenv(unsend.ENV_KEY_API_KEY)
	os.Setenv(unsend.ENV_KEY_BASE_URL, baseURL)
	defer os.Unsetenv(unsend.ENV_KEY_BASE_URL)

	client, err := unsend.NewClient()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if client.ApiKey != "test-api-key" {
		t.Errorf("expected apiKey to be 'test-api-key', got %s", client.ApiKey)
	}

	if client.BaseUrl.String() != baseURL {
		t.Errorf("expected baseUrl to be %s, got %s", baseURL, client.BaseUrl.String())
	}
}

func TestNewRequest(t *testing.T) {
	client := &unsend.Client{
		BaseUrl: &url.URL{Scheme: "https", Host: "api.example.com"},
	}

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("expected method to be GET, got %s", req.Method)
	}

	expectedURL := "https://api.example.com/test"
	if req.URL.String() != expectedURL {
		t.Errorf("expected URL to be %s, got %s", expectedURL, req.URL.String())
	}
}

func TestExecute(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	req, _ := client.NewRequest("GET", "/", nil)
	var result map[string]interface{}
	err := client.Execute(req, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if success, ok := result["success"].(bool); !ok || !success {
		t.Errorf("expected success to be true, got %v", result["success"])
	}
}
