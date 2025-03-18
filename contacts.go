package unsend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Contacts interface {
	GetContact(ctx context.Context, request GetContactRequest) (*GetContactResponse, error)
	CreateContact(ctx context.Context, request CreateContactRequest) (*ContactIdResponse, error)
	UpsertContact(ctx context.Context, request UpsertContactRequest) (*ContactIdResponse, error)
	UpdateContact(ctx context.Context, request UpdateContactRequest) (*ContactIdResponse, error)
	DeleteContact(ctx context.Context, request DeleteContactRequest) (*DeleteContactResponse, error)
}

type ContactsImpl struct {
	Client *Client
}

type GetContactRequest struct {
	ContactBookId string
	ContactId     string
}

type CreateContactRequest struct {
	ContactBookId string                 `json:"-"`
	Email         string                 `json:"email"`
	FirstName     string                 `json:"firstName,omitempty"`
	LastName      string                 `json:"lastName,omitempty"`
	Properties    map[string]interface{} `json:"properties,omitempty"`
	Subscribed    bool                   `json:"subscribed,omitempty"`
}

type UpsertContactRequest struct {
	ContactBookId string                 `json:"-"`
	ContactId     string                 `json:"-"`
	Email         string                 `json:"email"`
	FirstName     string                 `json:"firstName,omitempty"`
	LastName      string                 `json:"lastName,omitempty"`
	Properties    map[string]interface{} `json:"properties,omitempty"`
	Subscribed    bool                   `json:"subscribed,omitempty"`
}

type UpdateContactRequest struct {
	ContactBookId string                 `json:"-"`
	ContactId     string                 `json:"-"`
	FirstName     string                 `json:"firstName,omitempty"`
	LastName      string                 `json:"lastName,omitempty"`
	Properties    map[string]interface{} `json:"properties,omitempty"`
	Subscribed    bool                   `json:"subscribed,omitempty"`
}

type DeleteContactRequest struct {
	ContactBookId string
	ContactId     string
}

type GetContactResponse struct {
	Id            string                 `json:"id"`
	FirstName     string                 `json:"firstName"`
	LastName      string                 `json:"lastName"`
	Email         string                 `json:"email"`
	Subscribed    bool                   `json:"subscribed"`
	Properties    map[string]interface{} `json:"properties"`
	ContactBookID string                 `json:"contactBookId"`
	CreatedAt     string                 `json:"createdAt"`
	UpdatedAt     string                 `json:"updatedAt"`
}

type ContactIdResponse struct {
	ContactId string `json:"contactId"`
}

type DeleteContactResponse struct {
	Success bool `json:"success"`
}

func (c *ContactsImpl) GetContact(ctx context.Context, request GetContactRequest) (*GetContactResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: GetContactRequest not valid; %v", err.Errors)
	}

	path := "api/v1/contactBooks/" + request.ContactBookId + "/contacts/" + request.ContactId

	req, err := c.Client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return &GetContactResponse{}, errors.New("[ERROR]: Failed to create Contact.Get request")
	}

	response := new(GetContactResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return &GetContactResponse{}, err
	}

	return response, nil
}

func (c *ContactsImpl) CreateContact(ctx context.Context, request CreateContactRequest) (*ContactIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: CreateContactRequest not valid; %v", err.Errors)
	}

	path := "api/v1/contactBooks/" + request.ContactBookId + "/contacts/"

	req, err := c.Client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return &ContactIdResponse{}, errors.New("[ERROR]: Failed to create Contact.Create request")
	}

	response := new(ContactIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return &ContactIdResponse{}, err
	}
	return response, nil
}

func (c *ContactsImpl) UpsertContact(ctx context.Context, request UpsertContactRequest) (*ContactIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: UpsertContactRequest not valid; %v", err.Errors)
	}

	path := "api/v1/contactBooks/" + request.ContactBookId + "/contacts/" + request.ContactId

	req, err := c.Client.NewRequest(http.MethodPut, path, request)
	if err != nil {
		return &ContactIdResponse{}, errors.New("[ERROR]: Failed to create Contact.Upsert request")
	}

	response := new(ContactIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return &ContactIdResponse{}, err
	}
	return response, nil
}

func (c *ContactsImpl) UpdateContact(ctx context.Context, request UpdateContactRequest) (*ContactIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: UpdateContactRequest not valid; %v", err.Errors)
	}

	path := "api/v1/contactBooks/" + request.ContactBookId + "/contacts/" + request.ContactId

	req, err := c.Client.NewRequest(http.MethodPatch, path, request)
	if err != nil {
		return &ContactIdResponse{}, errors.New("[ERROR]: Failed to create Contact.Update request")
	}

	response := new(ContactIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return &ContactIdResponse{}, err
	}
	return response, nil
}

func (c *ContactsImpl) DeleteContact(ctx context.Context, request DeleteContactRequest) (*DeleteContactResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: DeleteContactRequest not valid; %v", err.Errors)
	}

	path := "api/v1/contactBooks/" + request.ContactBookId + "/contacts/" + request.ContactId

	req, err := c.Client.NewRequest(http.MethodDelete, path, request)
	if err != nil {
		return &DeleteContactResponse{
			Success: false,
		}, errors.New("[ERROR]: Failed to create Contact.Delete request")
	}

	response := new(DeleteContactResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return &DeleteContactResponse{
			Success: false,
		}, err
	}
	return response, nil
}

func (req CreateContactRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateContactRequest
	return json.Marshal(&struct {
		Subscribed bool `json:"subscribed"`
		*Alias
	}{
		Subscribed: req.Subscribed,
		Alias:      (*Alias)(&req),
	})
}

func (req UpsertContactRequest) MarshalJSON() ([]byte, error) {
	type Alias UpsertContactRequest
	return json.Marshal(&struct {
		Subscribed bool `json:"subscribed"`
		*Alias
	}{
		Subscribed: req.Subscribed,
		Alias:      (*Alias)(&req),
	})
}

func (req UpdateContactRequest) MarshalJSON() ([]byte, error) {
	type Alias UpdateContactRequest
	return json.Marshal(&struct {
		Subscribed bool `json:"subscribed"`
		*Alias
	}{
		Subscribed: req.Subscribed,
		Alias:      (*Alias)(&req),
	})
}
