package zipnova

import (
	"fmt"
)

type TrackingSort string

const (
	TrackingNewest TrackingSort = "newest"
	TrackingOldest TrackingSort = "oldest"
)

func (c *Client) GetTracking(shipmentID int, sort TrackingSort) (*TrackingResponse, error) {
	path := fmt.Sprintf("/shipments/%d/tracking", shipmentID)

	httpReq, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	if sort == "" {
		sort = TrackingNewest
	}
	q := httpReq.URL.Query()
	q.Add("sort", string(sort))
	httpReq.URL.RawQuery = q.Encode()

	var resp TrackingResponse
	if err := c.do(httpReq, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
