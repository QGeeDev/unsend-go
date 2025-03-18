package unsend_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/QGeeDev/unsend-go"
)

func TestCreateContact(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Contacts = &unsend.ContactsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/contactBooks/book123/contacts/" && r.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"contactId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.CreateContactRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.CreateContactRequest{
				ContactBookId: "book123",
				Email:         "test@example.com",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Not found request",
			request: unsend.CreateContactRequest{
				ContactBookId: "invalidBook",
				Email:         "test@example.com",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
			},
			expectedID:     "",
			expectedErrMsg: "received non-2xx response: 404 - {\"error\": \"not found\"}",
		},
		{
			name: "Invalid request",
			request: unsend.CreateContactRequest{
				ContactBookId: "",
				Email:         "test@example.com",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: CreateContactRequest not valid; ['ContactBookId' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Contacts.CreateContact(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.ContactId != tt.expectedID {
				t.Errorf("expected contactId to be '%s', got '%s'", tt.expectedID, response.ContactId)
			}
		})
	}
}

func TestGetContact(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Contacts = &unsend.ContactsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/contactBooks/book123/contacts/12345" && r.Method == http.MethodGet {
			response := unsend.GetContactResponse{
				Id:            "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Email:         "test@example.com",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
				ContactBookID: "book123",
				CreatedAt:     "2021-01-01T00:00:00Z",
				UpdatedAt:     "2021-01-01T00:00:00Z",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	tests := []struct {
		name            string
		request         unsend.GetContactRequest
		expectedContact *unsend.GetContactResponse
		expectedErrMsg  string
	}{
		{
			name: "Valid request",
			request: unsend.GetContactRequest{
				ContactBookId: "book123",
				ContactId:     "12345",
			},
			expectedContact: nil,
			expectedErrMsg:  "",
		},
		{
			name: "Not found request",
			request: unsend.GetContactRequest{
				ContactBookId: "invalidBook",
				ContactId:     "54321",
			},
			expectedContact: nil,
			expectedErrMsg:  "received non-2xx response: 404 - {\"error\": \"not found\"}",
		},
		{
			name: "Invalid request",
			request: unsend.GetContactRequest{
				ContactBookId: "",
				ContactId:     "12345",
			},
			expectedContact: nil,
			expectedErrMsg:  "[ERROR]: GetContactRequest not valid; ['ContactBookId' is required]",
		},
	}

	client.BaseUrl, _ = url.Parse(server.URL)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Contacts.GetContact(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && tt.expectedContact != nil {
				assertContactResponseEqual(t, tt.expectedContact, response)
			}
		})
	}
}

func TestUpdateContact(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Contacts = &unsend.ContactsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/contactBooks/book123/contacts/12345" && r.Method == http.MethodPatch {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"contactId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.UpdateContactRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.UpdateContactRequest{
				ContactBookId: "book123",
				ContactId:     "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Not found request",
			request: unsend.UpdateContactRequest{
				ContactBookId: "book123",
				ContactId:     "54321",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "received non-2xx response: 404 - {\"error\": \"not found\"}",
		},
		{
			name: "Invalid request",
			request: unsend.UpdateContactRequest{
				ContactBookId: "",
				ContactId:     "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpdateContactRequest not valid; ['ContactBookId' is required]",
		},
		{
			name: "Multiple Invalid Field request",
			request: unsend.UpdateContactRequest{
				ContactBookId: "",
				ContactId:     "",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpdateContactRequest not valid; ['ContactBookId' is required 'ContactId' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Contacts.UpdateContact(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.ContactId != tt.expectedID {
				t.Errorf("expected contactId to be '%s', got '%s'", tt.expectedID, response.ContactId)
			}
		})
	}
}

func TestUpsertContact(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Contacts = &unsend.ContactsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/contactBooks/book123/contacts/12345" && r.Method == http.MethodPut {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"contactId": "12345"}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.UpsertContactRequest
		expectedID     string
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.UpsertContactRequest{
				ContactBookId: "book123",
				ContactId:     "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Subscribed:    true,
				Email:         "bill.gates@microsoft.com",
				Properties:    map[string]interface{}{},
			},
			expectedID:     "12345",
			expectedErrMsg: "",
		},
		{
			name: "Not found request",
			request: unsend.UpsertContactRequest{
				ContactBookId: "book123",
				ContactId:     "54321",
				FirstName:     "John",
				LastName:      "Doe",
				Email:         "bill.gates@microsoft.com",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "received non-2xx response: 404 - {\"error\": \"not found\"}",
		},
		{
			name: "Invalid request",
			request: unsend.UpsertContactRequest{
				ContactBookId: "",
				ContactId:     "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Email:         "bill.gates@microsoft.com",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpsertContactRequest not valid; ['ContactBookId' is required]",
		},
		{
			name: "Multiple Invalid Field request",
			request: unsend.UpsertContactRequest{
				ContactBookId: "",
				ContactId:     "12345",
				FirstName:     "John",
				LastName:      "Doe",
				Email:         "",
				Subscribed:    true,
				Properties:    map[string]interface{}{},
			},
			expectedID:     "",
			expectedErrMsg: "[ERROR]: UpsertContactRequest not valid; ['ContactBookId' is required 'Email' is required]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Contacts.UpsertContact(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil && tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.ContactId != tt.expectedID {
				t.Errorf("expected contactId to be '%s', got '%s'", tt.expectedID, response.ContactId)
			}
		})
	}
}

func TestDeleteContact(t *testing.T) {
	client := &unsend.Client{
		Client: &http.Client{},
	}

	client.Contacts = &unsend.ContactsImpl{Client: client}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/contactBooks/book123/contacts/12345" && r.Method == http.MethodDelete {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer server.Close()

	client.BaseUrl, _ = url.Parse(server.URL)

	tests := []struct {
		name           string
		request        unsend.DeleteContactRequest
		expectedErrMsg string
	}{
		{
			name: "Valid request",
			request: unsend.DeleteContactRequest{
				ContactBookId: "book123",
				ContactId:     "12345",
			},
			expectedErrMsg: "",
		},
		{
			name: "Not found request",
			request: unsend.DeleteContactRequest{
				ContactBookId: "book123",
				ContactId:     "54321",
			},
			expectedErrMsg: "received non-2xx response: 404 - {\"error\": \"not found\"}",
		},
		// {
		// 	name: "Invalid request",
		// 	request: unsend.DeleteContactRequest{
		// 		ContactBookId: "",
		// 		ContactId:     "12345",
		// 	},
		// 	expectedErrMsg: "[ERROR]: DeleteContactRequest not valid; ['ContactBookId' is required]",
		// },
		// {
		// 	name: "Multiple Invalid Field request",
		// 	request: unsend.DeleteContactRequest{
		// 		ContactBookId: "",
		// 		ContactId:     "",
		// 	},
		// 	expectedErrMsg: "[ERROR]: DeleteContactRequest not valid; ['ContactBookId' is required 'ContactId' is required]",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			response, err := client.Contacts.DeleteContact(ctx, tt.request)
			if err != nil && tt.expectedErrMsg == "" {
				t.Fatalf("expected no error, got %v", err)
			}
			if err == nil &&
				tt.expectedErrMsg != "" {
				t.Fatalf("expected error %v, got no error", tt.expectedErrMsg)
			}
			if err != nil && tt.expectedErrMsg != "" && err.Error() != tt.expectedErrMsg {
				t.Fatalf("expected error %v, got %v", tt.expectedErrMsg, err)
			}
			if response != nil && response.Success != true && tt.expectedErrMsg == ""{
				t.Errorf("expected success to be true, got %v", response.Success)
			}
		})
	}
}

func assertContactResponseEqual(t *testing.T, expected, actual *unsend.GetContactResponse) {
	if expected == nil || actual == nil {
		if expected != actual {
			t.Errorf("expected contact to be '%v', got '%v'", expected, actual)
		}
		return
	}

	if expected.Id != actual.Id ||
		expected.FirstName != actual.FirstName ||
		expected.LastName != actual.LastName ||
		expected.Email != actual.Email ||
		expected.Subscribed != actual.Subscribed {
		t.Errorf("expected contact to be '%v', got '%v'", *expected, *actual)
	}
}
