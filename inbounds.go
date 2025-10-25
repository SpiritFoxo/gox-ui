package api

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

// GetClientByEmail gets most of the client info
func (i *Inbound) GetClientByEmail(api *Api, email string) *Client {
	for _, client := range i.Settings.Value.Clients {
		if client.Value.Email == email {
			return &Client{
				Comment:    client.Value.Comment,
				Email:      client.Value.Email,
				Enable:     client.Value.Enable,
				ExpiryTime: client.Value.ExpiryTime,
				Flow:       client.Value.Flow,
				UUID:       client.Value.UUID,
				LimitIP:    client.Value.LimitIP,
				Reset:      client.Value.Reset,
				SubId:      client.Value.SubId,
				TgId:       client.Value.TgId,
				TotalGB:    client.Value.TotalGB,
				InboundId:  client.Value.InboundId,
				Up:         client.Value.Up,
				Down:       client.Value.Down,
				AllTime:    client.Value.AllTime,
			}
		}
	}
	return nil
}
