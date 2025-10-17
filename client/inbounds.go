package client

import (
	"context"
	"fmt"
)

// ListInbounds gets array of all available inbounds from the server.
func (a *Api) ListInbounds(ctx context.Context) (*ListInboundsResponse, error) {
	var resp ListInboundsResponse
	return &resp, a.DoRequest(ctx, "GET", "/inbounds/list", nil, &resp)
}

// GetInbound gets inbound by provided id.
func (a *Api) GetInbound(ctx context.Context, id uint) (*GetInboundResponse, error) {
	var resp GetInboundResponse

	endpoint := fmt.Sprintf("/inbounds/get/%d", id)
	fmt.Println(endpoint)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}
