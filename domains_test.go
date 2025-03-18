package unsend_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/QGeeDev/unsend-go"
)

func TestGetDomains(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Domains = &unsend.DomainsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/domains" && r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[
				{
					"id": 1,
					"name": "unsend.dev",
					"teamId": 1,
					"status": "NOT_STARTED",
					"region": "us-east-1",
					"clickTracking": false,
					"openTracking": false,
					"publicKey": "key123",
					"dkimStatus": "SUCCESS",
					"spfDetails": "SUCCESS",
					"createdAt": "2025-01-01T00:00:00Z",
					"updatedAt": "2025-01-01T00:00:00Z"
				}
			]`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	tests := []struct {
		name            string
		expectedDomains *[]unsend.GetDomainsResponse
		expectedErrMsg  string
	}{
		{
			name: "Response Unmarshals correctly",
			expectedDomains: &[]unsend.GetDomainsResponse{
				{
					Id:            1,
					Name:          "unsend.dev",
					TeamId:        1,
					Status:        "NOT_STARTED",
					Region:        "us-east-1",
					ClickTracking: false,
					OpenTracking:  false,
					PublicKey:     "key123",
					DkimStatus:    "SUCCESS",
					SpfDetails:    "SUCCESS",
					CreatedAt:     "2025-01-01T00:00:00Z",
					UpdatedAt:     "2025-01-01T00:00:00Z",
				},
			},
			expectedErrMsg: "",
		},
	}

	client.BaseUrl, _ = url.Parse(server.URL)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Domains.GetDomains(ctx)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if !reflect.DeepEqual(response, tt.expectedDomains) {
				t.Errorf("expected response to be %v, got %v", tt.expectedDomains, response)
			}
		})
	}
}
