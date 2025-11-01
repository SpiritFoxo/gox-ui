package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetClientTrafficByEmail returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByEmail(ctx context.Context, client *Client) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTraffics/%s", client.Email)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetClientTrafficByUUID returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByUUID(ctx context.Context, client *Client) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTrafficsById/%s", client.UUID)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// AddClient allows to create new client inside selected inbound
func (a *Api) AddClient(ctx context.Context, inbound *Inbound, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/addClient"
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: *client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(inbound.ID), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// ResetClientTraffic resets client traffic
func (a *Api) ResetClientTraffic(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/resetClientTraffic/%s", *client.InboundId, client.Email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// UpdateClientInfo updates information about client
func (a *Api) UpdateClientinfo(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/updateClient/%s", client.UUID)
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: *client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(*client.InboundId), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// GetClientIpAdress returns clients`s IP adress
func (a *Api) GetClientIpAdress(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clientIps/%s", client.Email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// ClearClientIps clears IP assigned to a client
func (a *Api) ClearClientIps(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clearClientIps/%s", client.Email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteClient deletes client from inbound
func (a *Api) DeleteClient(ctx context.Context, client *Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/delClient/%s", *client.InboundId, client.UUID)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}
