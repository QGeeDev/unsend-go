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

func TestSendEmail(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Emails = &unsend.EmailsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/emails" && r.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"emailId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.SendEmailRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.SendEmailRequest{
				To:         []string{"a@b.c"},
				From:       "test@unsend.dev",
				Subject:    "Test email",
				TemplateId: "12345",
				Text:       "Hello, World!",
				ReplyTo:    []string{"replyto@unsend.dev"},
				ScheduleAt: "2021-01-01T00:00:00Z",
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Invalid request",
			request: unsend.SendEmailRequest{
				To:         []string{},
				From:       "test@unsend.dev",
				Subject:    "Test email",
				TemplateId: "12345",
				Text:       "Hello, World!",
				ReplyTo:    []string{"replyto@unsend.dev"},
				ScheduleAt: "2021-01-01T00:00:00Z",
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: SendEmailRequest not valid; ['To' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Emails.SendEmail(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.EmailId != tt.expectedID {
				t.Errorf("expected email ID to be %s, got %s", tt.expectedID, response.EmailId)
			}
		})
	}
}

func TestGetEmail(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Emails = &unsend.EmailsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/emails/12345" && r.Method == http.MethodGet {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
			"id": "12345",
			"teamId": 123,
    		"to": ["recipient@unsend.dev"],
    		"from": "sender@unsend.dev",
    		"subject": "Test Email",
			"text": "Hello, World!",
			"html": "<p>Hello, World!</p>",
			"createdAt": "2021-01-01T00:00:00Z",
			"updatedAt": "2021-01-01T00:00:00Z",
			"emailEvents": [{"emailId": "12345", "status": "SENT", "createdAt": "2021-01-01T00:00:00Z", "data": {"timestamp": "2021-01-01T00:00:00Z"}}],
    		"replyTo": ["replyto@unsend.dev"],
    		"cc": ["cc@unsend.dev"],
    		"bcc": ["bcc@unsend.dev"]}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name             string
		requestId        unsend.GetEmailRequest
		expectedResponse unsend.GetEmailResponse
		expectedErrMsg   string
	}{
		{
			name: "Valid request",
			requestId: unsend.GetEmailRequest{
				EmailId: "12345",
			},
			expectedResponse: unsend.GetEmailResponse{
				Id:        "12345",
				TeamId:    123,
				To:        []string{"recipient@unsend.dev"},
				From:      "sender@unsend.dev",
				Subject:   "Test Email",
				Html:      "<p>Hello, World!</p>",
				Text:      "Hello, World!",
				CreatedAt: "2021-01-01T00:00:00Z",
				UpdatedAt: "2021-01-01T00:00:00Z",
				EmailEvents: []unsend.EmailEvents{
					{
						EmailId:   "12345",
						Status:    "SENT",
						CreatedAt: "2021-01-01T00:00:00Z",
						Data:      map[string]interface{}{"timestamp": "2021-01-01T00:00:00Z"},
					},
				},
				ReplyTo: []string{"replyto@unsend.dev"},
				Cc:      []string{"cc@unsend.dev"},
				Bcc:     []string{"bcc@unsend.dev"},
			},
			expectedErrMsg: "",
		},
		{
			name: "Not found request",
			requestId: unsend.GetEmailRequest{
				EmailId: "54321",
			},
			expectedResponse: unsend.GetEmailResponse{},
			expectedErrMsg:   `received non-2xx response: 404 - {"error": "not found"}`,
		},
		{
			name: "Invalid request",
			requestId: unsend.GetEmailRequest{
				EmailId: "",
			},
			expectedResponse: unsend.GetEmailResponse{},
			expectedErrMsg:   `[ERROR]: GetEmailRequest not valid; ['EmailId' is required]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Emails.GetEmail(ctx, tt.requestId)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if tt.expectedErrMsg == "" && !reflect.DeepEqual(response, &tt.expectedResponse) {
				t.Errorf("expected response to be %v, got %v", tt.expectedResponse, response)
			}
		})
	}
}

func TestUpdateSchedule(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Emails = &unsend.EmailsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/emails/12345" && r.Method == http.MethodPatch {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"emailId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.UpdateScheduleRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.UpdateScheduleRequest{
				EmailId:     "12345",
				ScheduledAt: "2021-01-01T00:00:00Z",
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Invalid request - missing email ID",
			request: unsend.UpdateScheduleRequest{
				EmailId:     "",
				ScheduledAt: "2021-01-01T00:00:00Z",
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpdateScheduleRequest not valid; ['EmailId' is required]",
		},
		{
			name: "Invalid request - missing scheduled at",
			request: unsend.UpdateScheduleRequest{
				EmailId:     "12345",
				ScheduledAt: "",
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpdateScheduleRequest not valid; ['ScheduledAt' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Emails.UpdateSchedule(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.EmailId != tt.expectedID {
				t.Errorf("expected email ID to be %s, got %s", tt.expectedID, response.EmailId)
			}
		})
	}
}

func TestCancelSchedule(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Emails = &unsend.EmailsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/emails/12345/cancel" && r.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"emailId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.CancelScheduleRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.CancelScheduleRequest{
				EmailId:     "12345",
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Invalid request - missing email ID",
			request: unsend.CancelScheduleRequest{
				EmailId:     "",
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: CancelScheduleRequest not valid; ['EmailId' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Emails.CancelSchedule(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.EmailId != tt.expectedID {
				t.Errorf("expected email ID to be %s, got %s", tt.expectedID, response.EmailId)
			}
		})
	}
}