package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetClientTrafficByEmail returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByEmail(ctx context.Context, email string) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTraffics/%s", email)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// GetClientTrafficByUUID returns amount of downloaded and uploaded data by a client
func (a *Api) GetClientTrafficByUUID(ctx context.Context, uuid string) (*GetClientTrafficResponse, error) {
	var resp GetClientTrafficResponse
	endpoint := fmt.Sprintf("/inbounds/getClientTrafficsById/%s", uuid)
	return &resp, a.DoRequest(ctx, "GET", endpoint, nil, &resp)
}

// AddClient allows to create new client inside selected inbound
func (a *Api) AddClient(ctx context.Context, inboundId uint, client Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := "/inbounds/addClient"
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(inboundId), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// ResetClientTraffic resets client traffic
func (a *Api) ResetClientTraffic(ctx context.Context, inboundId uint, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/resetClientTraffic/%s", inboundId, email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// UpdateClientInfo updates information about client
func (a *Api) UpdateClientinfo(ctx context.Context, inboundId uint, client Client) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/updateClient/%s", client.UUID)
	settings := Settings{
		Clients: []JSONString[Client]{{
			Value: client,
		}},
	}
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %w", err)
	}

	formData := map[string]string{
		"id":       strconv.FormatUint(uint64(inboundId), 10),
		"settings": string(settingsJSON),
	}

	return &resp, a.DoFormDataRequest(ctx, "POST", endpoint, formData, &resp)
}

// GetClientIpAdress returns clients`s IP adress
func (a *Api) GetClientIpAdress(ctx context.Context, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clientIps/%s", email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// ClearClientIps clears IP assigned to a client
func (a *Api) ClearClientIps(ctx context.Context, email string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/clearClientIps/%s", email)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}

// DeleteClient deletes client from inbound
func (a *Api) DeleteClient(ctx context.Context, inboundId uint, uuid string) (*MessageResponse, error) {
	var resp MessageResponse
	endpoint := fmt.Sprintf("/inbounds/%d/delClient/%s", inboundId, uuid)
	return &resp, a.DoRequest(ctx, "POST", endpoint, nil, &resp)
}
