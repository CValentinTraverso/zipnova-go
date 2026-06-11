package zipnova

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateShipmentRequest struct {
	AccountID           int                  `json:"account_id"`
	ExternalID          string               `json:"external_id"`
	ServiceType         string               `json:"service_type"`
	LogisticType        string               `json:"logistic_type,omitempty"`
	CarrierID           int                  `json:"carrier_id,omitempty"`
	SortBy              string               `json:"sort_by,omitempty"`
	OriginID            string               `json:"origin_id"`
	DeclaredValue       float64              `json:"declared_value"`
	Source              string               `json:"source,omitempty"`
	TypePackaging       string               `json:"type_packaging,omitempty"`
	ProcessImmediately  int                  `json:"process_immediately,omitempty"`
	Destination         ShipmentDestination  `json:"destination"`
	Packages            []Package            `json:"packages,omitempty"`
	Items               []Item               `json:"items,omitempty"`
}

func (c *Client) CreateShipment(req *CreateShipmentRequest) (*ShipmentResource, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := c.newRequest("POST", "/shipments", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var resp ShipmentResource
	if err := c.do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetShipment(id int) (*ShipmentResource, error) {
	path := fmt.Sprintf("/shipments/%d", id)

	httpReq, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var resp ShipmentResource
	if err := c.do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
