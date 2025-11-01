package api

import (
	"context"
	"fmt"
)

// ListInbounds gets array of all available inbounds from the server
func (a *Api) ListInbounds(ctx context.Context) (*[]Inbound, error) {
	var resp ListInboundsResponse
	endpoint := "inbounds/list"
	return &resp.Obj, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetInbound gets inbound by provided id
func (a *Api) GetInbound(ctx context.Context, inboundId uint) (*Inbound, error) {
	var resp GetInboundResponse
	endpoint := fmt.Sprintf("/inbounds/get/%d", inboundId)
	return &resp.Obj, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// ReserAllTraffic is used to reset the traffic statistics for all inbounds within the system
func (a *Api) ResetAllTraffic(ctx context.Context) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/resetAllTraffics"
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// ResetTraffic is used to reset the traffic statistics for all clients associated with a specific inbound
func (a *Api) ResetTraffic(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/resetAllClientTraffics/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteInbound deletes inbound from server
func (a *Api) DeleteInbound(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	enpoint := fmt.Sprintf("/inbounds/del/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", enpoint, nil, &resp)
}

// DeleteDepletedClients delets all depleted clients from inbound
func (a *Api) DeleteDepletedClients(ctx context.Context, inboundId uint) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("inbounds/delDepletedClients/%d", inboundId)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}
