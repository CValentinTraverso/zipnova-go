package zipnova

import (
	"bytes"
	"encoding/json"
)

type QuoteRequest struct {
	AccountID          int         `json:"account_id"`
	OriginID           int         `json:"origin_id,omitempty"`
	Source             string      `json:"source"`
	Destination        Destination `json:"destination"`
	DeclaredValue      float64     `json:"declared_value"`
	Items              []Item      `json:"items,omitempty"`
	Packages           []Package   `json:"packages,omitempty"`
	TypePackaging      string      `json:"type_packaging,omitempty"`
	LogisticType       string      `json:"logistic_type,omitempty"`
	ServiceType        string      `json:"service_type,omitempty"`
	SortBy             string      `json:"sort_by,omitempty"`
	AvoidRules         bool        `json:"avoid_rules,omitempty"`
	IncludeDropoffPts  int         `json:"include_dropoff_points,omitempty"`
}

func (c *Client) Quote(req *QuoteRequest) (*QuoteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := c.newRequest("POST", "/shipments/quote", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var resp QuoteResponse
	if err := c.do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
