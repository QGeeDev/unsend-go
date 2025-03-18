package unsend

import (
	"context"
	"net/http"
)

type Domains interface {
	GetDomains(ctx context.Context) (*[]GetDomainsResponse, error)
}

type GetDomainsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	TeamId        int    `json:"teamId"`
	Status        string `json:"status"`
	PublicKey     string `json:"publicKey"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Region        string `json:"region"`
	ClickTracking bool   `json:"clickTracking"`
	OpenTracking  bool   `json:"openTracking"`
	DkimStatus    string `json:"dkimStatus,omitempty"`
	SpfDetails    string `json:"spfDetails,omitempty"`
}

type DomainsImpl struct {
	Client *Client
}

func (d *DomainsImpl) GetDomains(ctx context.Context) (*[]GetDomainsResponse, error) {
	path := "api/v1/domains"

	req, err := d.Client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	response := new([]GetDomainsResponse)
	err = d.Client.Execute(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
