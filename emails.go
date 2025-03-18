package unsend

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Emails interface {
	GetEmail(ctx context.Context, request GetEmailRequest) (*GetEmailResponse, error)
	SendEmail(ctx context.Context, request SendEmailRequest) (*EmailIdResponse, error)
	UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*EmailIdResponse, error)
	CancelSchedule(ctx context.Context, request CancelScheduleRequest) (*EmailIdResponse, error)
}

type EmailEvents struct {
	EmailId   string      `json:"emailId"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"createdAt"`
	Data      interface{} `json:"data"`
}

type Attachments struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type GetEmailRequest struct {
	EmailId string
}

type GetEmailResponse struct {
	Id          string        `json:"id"`
	TeamId      int           `json:"teamId"`
	To          []string      `json:"to"`
	From        string        `json:"from"`
	Subject     string        `json:"subject"`
	Html        string        `json:"html"`
	Text        string        `json:"text"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
	EmailEvents []EmailEvents `json:"emailEvents"`
	ReplyTo     []string      `json:"replyTo"`
	Cc          []string      `json:"cc"`
	Bcc         []string      `json:"bcc"`
}

type SendEmailRequest struct {
	To          []string               `json:"to"`
	From        string                 `json:"from"`
	Subject     string                 `json:"subject,omitempty"`
	TemplateId  string                 `json:"templateId,omitempty"`
	Variables   map[string]interface{} `json:"variables,omitempty"`
	ReplyTo     []string               `json:"replyTo,omitempty"`
	Cc          []string               `json:"cc,omitempty"`
	Bcc         []string               `json:"bcc,omitempty"`
	Text        string                 `json:"text,omitempty"`
	Html        string                 `json:"html,omitempty"`
	Attachments []Attachments          `json:"attachments,omitempty"`
	ScheduleAt  string                 `json:"scheduleAt,omitempty"`
}

type EmailIdResponse struct {
	EmailId string `json:"emailId"`
}

type UpdateScheduleRequest struct {
	EmailId     string `json:"-"`
	ScheduledAt string `json:"scheduledAt"`
}

type CancelScheduleRequest struct {
	EmailId string
}

type EmailsImpl struct {
	Client *Client
}

func (e *EmailsImpl) GetEmail(ctx context.Context, request GetEmailRequest) (*GetEmailResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: GetEmailRequest not valid; %v", err.Errors)
	}

	path := "api/v1/emails/" + request.EmailId

	req, err := e.Client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.New("[ERROR]: Failed to create Email.Get request")
	}

	response := new(GetEmailResponse)
	err = e.Client.Execute(req, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *EmailsImpl) SendEmail(ctx context.Context, request SendEmailRequest) (*EmailIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: SendEmailRequest not valid; %v", err.Errors)
	}

	path := "api/v1/emails"

	req, err := c.Client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, errors.New("[ERROR]: Failed to create Send Email request")
	}

	response := new(EmailIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *EmailsImpl) UpdateSchedule(ctx context.Context, request UpdateScheduleRequest) (*EmailIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: UpdateScheduleRequest not valid; %v", err.Errors)
	}

	path := "api/v1/emails/" + request.EmailId

	req, err := c.Client.NewRequest(http.MethodPatch, path, request)
	if err != nil {
		return nil, errors.New("[ERROR]: Failed to create Update Schedule request")
	}

	response := new(EmailIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *EmailsImpl) CancelSchedule(ctx context.Context, request CancelScheduleRequest) (*EmailIdResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("[ERROR]: CancelScheduleRequest not valid; %v", err.Errors)
	}

	path := "api/v1/emails/" + request.EmailId + "/cancel"

	req, err := c.Client.NewRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, errors.New("[ERROR]: Failed to create Update Schedule request")
	}

	response := new(EmailIdResponse)
	err = c.Client.Execute(req, response)

	if err != nil {
		return nil, err
	}
	return response, nil
}
